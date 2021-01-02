package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/t-tiger/survey/server/config"

	"github.com/t-tiger/survey/server/entity"

	"github.com/prometheus/common/log"

	"github.com/t-tiger/survey/server/cerrors"
)

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())

	reason := cerrors.GetReason(err)
	switch reason {
	case cerrors.Duplicated:
		http.Error(w, err.Error(), http.StatusConflict)
	case cerrors.ValidationFailed:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func createToken(user entity.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	s, err := t.SignedString([]byte(config.Config.SecretKey))
	if err != nil {
		return "", cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	return s, nil
}
