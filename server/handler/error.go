package handler

import (
	"net/http"

	"github.com/prometheus/common/log"

	"github.com/t-tiger/survey/server/cerrors"
)

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())

	code := cerrors.GetCode(err)
	switch code {
	case cerrors.Duplicated:
		http.Error(w, err.Error(), http.StatusConflict)
	case cerrors.ValidationFailed:
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
