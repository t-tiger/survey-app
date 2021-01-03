package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	repo repository.User
}

func NewUserAuth(repo repository.User) *UserAuth {
	return &UserAuth{repo: repo}
}

func (u *UserAuth) Call(ctx context.Context, email, password string) (entity.User, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return entity.User{}, cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	if user == nil {
		return entity.User{}, cerrors.Errorf(cerrors.Duplicated, "email has not been registered")
	}
	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return entity.User{}, cerrors.Errorf(cerrors.InvalidInput, "password is wrong")
	}
	return *user, nil
}
