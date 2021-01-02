package usecase

import (
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{repo: repo}
}

func (u *User) Create(email, password string) (entity.User, error) {
	panic("implement me")
}
