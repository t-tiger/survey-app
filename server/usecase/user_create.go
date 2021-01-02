package usecase

import (
	"context"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("^[a-zA-Z0-9.@!#$%&'*+/=?^_`{|}~-]+$")

type UserCreate struct {
	repo repository.User
}

func NewUserCreate(repo repository.User) *UserCreate {
	return &UserCreate{repo: repo}
}

func (u *UserCreate) Call(ctx context.Context, name, email, password string) (user entity.User, err error) {
	// validate argument
	if len(name) == 0 {
		return user, cerrors.Errorf(cerrors.ValidationFailed, "name must not be empty")
	}
	if len(password) < 5 {
		return user, cerrors.Errorf(cerrors.ValidationFailed, "password length must be greater than or equal to 5")
	}
	if !passwordRegex.MatchString(password) {
		return user, cerrors.Errorf(cerrors.ValidationFailed, "password format is invalid")
	}
	if !emailRegex.MatchString(email) {
		return user, cerrors.Errorf(cerrors.ValidationFailed, "email format is invalid")
	}
	duplicated, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return user, cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	if duplicated != nil {
		return user, cerrors.Errorf(cerrors.Duplicated, "email has already been registered")
	}

	// create digested password
	digestPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return user, cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	// persist user
	return u.repo.Create(ctx, name, email, string(digestPass))
}
