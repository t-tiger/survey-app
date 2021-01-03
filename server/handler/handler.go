package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/config"
)

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	message := err.Error()
	reason := cerrors.GetReason(err)
	switch reason {
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
