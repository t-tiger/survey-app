package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/usecase"
)

type User struct {
	createUsecase *usecase.UserCreate
}

func NewUser(createUsecase *usecase.UserCreate) *User {
	return &User{createUsecase: createUsecase}
}

type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createResponse struct {
	Token string `json:"token"`
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "", 500)
		return
	}
	user, err := h.createUsecase.Call(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		handleError(err, w)
		return
	}
	token, err := createToken(user)
	if err != nil {
		handleError(err, w)
		return
	}
	res := createResponse{Token: token}
	render.JSON(w, r, &res)
}
