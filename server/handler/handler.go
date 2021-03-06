package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/prometheus/common/log"
	"github.com/t-tiger/survey/server/cerrors"
)

// decoder declares globally for the sake of caching
var decoder = schema.NewDecoder()

type errResponse struct {
	Message string `json:"message"`
}

func handleError(err error, w http.ResponseWriter) {
	log.Error(err.Error())
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	message := err.Error()
	reason := cerrors.GetReason(err)
	switch reason {
	case cerrors.InvalidInput:
		w.WriteHeader(http.StatusBadRequest)
	case cerrors.Unauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case cerrors.Forbidden:
		w.WriteHeader(http.StatusForbidden)
	case cerrors.NotFound:
		message = "not found"
		w.WriteHeader(http.StatusNotFound)
	case cerrors.Duplicated:
		w.WriteHeader(http.StatusConflict)
	default:
		message = "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
	}

	res := errResponse{Message: message}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Errorf("failed to encode error message: %+v", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(message))
	}
}
