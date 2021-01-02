package handler

import (
	"net/http"

	"github.com/t-tiger/survey/server/usecase"
)

type User struct {
	userUsecase *usecase.User
}

func NewUser(userUsecase *usecase.User) *User {
	return &User{userUsecase: userUsecase}
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
