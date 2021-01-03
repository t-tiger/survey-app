package usecase

import (
	"context"
	"regexp"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
	"golang.org/x/crypto/bcrypt"
)

const passWordDigestCost = 10

var (
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	passwordRegex = regexp.MustCompile("^[a-zA-Z0-9.@!#$%&'*+/=?^_`{|}~-]+$")
)

type UserCreate struct {
	repo repository.User
}

func NewUserCreate(repo repository.User) *UserCreate {
	return &UserCreate{repo: repo}
}

func (u *UserCreate) Call(ctx context.Context, name, email, password string) (user entity.User, err error) {
	// validate argument
	if len(name) == 0 {
		return user, cerrors.Errorf(cerrors.InvalidInput, "name must not be empty")
	}
	if len(password) < 5 {
		return user, cerrors.Errorf(cerrors.InvalidInput, "password length must be greater than or equal to 5")
	}
	if !passwordRegex.MatchString(password) {
		return user, cerrors.Errorf(cerrors.InvalidInput, "password format is invalid")
	}
	if !emailRegex.MatchString(email) {
		return user, cerrors.Errorf(cerrors.InvalidInput, "email format is invalid")
	}
	duplicated, err := u.repo.FindBy(ctx, email)
	if err != nil {
		return user, err
	}
	if duplicated != nil {
		return user, cerrors.Errorf(cerrors.Duplicated, "email has already been registered")
	}

	// create digested password
	digestPass, err := bcrypt.GenerateFromPassword([]byte(password), passWordDigestCost)
	if err != nil {
		return user, cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	// persist user
	return u.repo.Create(ctx, name, email, string(digestPass))
}
