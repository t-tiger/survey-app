package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type UserFind struct {
	repo repository.User
}

func NewUserFind(repo repository.User) *UserFind {
	return &UserFind{repo: repo}
}

func (u *UserFind) Call(ctx context.Context, id string) (entity.User, error) {
	return u.repo.FindBy(ctx, id)
}
