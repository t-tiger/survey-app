package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/config"
)

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())

	reason := cerrors.GetReason(err)
	switch reason {
	case cerrors.Duplicated:
		http.Error(w, err.Error(), http.StatusConflict)
	case cerrors.InvalidInput:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
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
