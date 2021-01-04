package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/config"
)

const (
	ctxUserID       = "userID"
	tokenCookieName = "_survey_app_token"
)

// createToken generates json web token
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
	cookie, err := r.Cookie(tokenCookieName)
	if err != nil {
		return nil
	}
	tokenStr := cookie.Value
	if len(tokenStr) == 0 {
		return nil
	}

	// decode json web token
	c := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenStr, c, func(t *jwt.Token) (interface{}, error) {
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
