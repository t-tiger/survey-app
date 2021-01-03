package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type User struct {
	authUsecase   *usecase.UserAuth
	createUsecase *usecase.UserCreate
}

func NewUser(authUsecase *usecase.UserAuth, createUsecase *usecase.UserCreate) *User {
	return &User{
		authUsecase:   authUsecase,
		createUsecase: createUsecase,
	}
}

type userLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userLoginResponse struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}

func (h *User) Login(w http.ResponseWriter, r *http.Request) {
	var req userLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}
	user, err := h.authUsecase.Call(r.Context(), req.Email, req.Password)
	if err != nil {
		handleError(err, w)
		return
	}
	token, err := createToken(user.ID)
	if err != nil {
		handleError(err, w)
		return
	}
	res := userLoginResponse{Token: token, User: user}
	render.JSON(w, r, &res)
}

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userCreateResponse = userLoginResponse

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var req userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}
	user, err := h.createUsecase.Call(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		handleError(err, w)
		return
	}
	token, err := createToken(user.ID)
	if err != nil {
		handleError(err, w)
		return
	}
	res := userCreateResponse{Token: token, User: user}
	render.JSON(w, r, &res)
}
