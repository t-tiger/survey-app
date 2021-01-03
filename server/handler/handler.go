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
