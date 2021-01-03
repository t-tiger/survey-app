package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/prometheus/common/log"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/config"
)

const ctxUserID = "userID"

// decoder declares globally for the sake of caching
var decoder = schema.NewDecoder()

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	message := err.Error()
	reason := cerrors.GetReason(err)
	switch reason {
	case cerrors.Unauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case cerrors.Duplicated:
		w.WriteHeader(http.StatusConflict)
	case cerrors.InvalidInput:
		w.WriteHeader(http.StatusUnprocessableEntity)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		message = "internal server error" // hide actual error message
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"message": message}); err != nil {
		log.Errorf("failed to encode error message: %+v", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(message))
	}
}

func createToken(userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	s, err := t.SignedString([]byte(config.Config.SecretKey))
	if err != nil {
		return "", cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	return s, nil
}

// retrieveUserID returns userID by decoding jwt with Authorization value
func retrieveUserID(r *http.Request) *string {
	authHead := r.Header.Get("Authorization")
	if len(authHead) == 0 {
		return nil
	}
	tokenStr := strings.Replace(authHead, "Bearer ", "", 1)

	// decode json web token
	c := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, c, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.SecretKey), nil
	})
	if err != nil {
		return nil
	}
	if v, ok := c["user_id"].(string); ok {
		return &v
	}
	return nil
}

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := retrieveUserID(r)
		if userID == nil {
			handleError(cerrors.Errorf(cerrors.Unauthorized, "unauthorized user"), w)
			return
		}
		// set authorized user's id in context
		ctx := context.WithValue(r.Context(), ctxUserID, *userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
