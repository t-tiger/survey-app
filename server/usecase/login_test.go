package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin_Call(t *testing.T) {
	// initialize appropriate password
	pw := []byte("foobar")
	pwDigest, err := bcrypt.GenerateFromPassword(pw, passWordDigestCost)
	require.NoError(t, err)

	// mock repository
	repo := &userRepoMock{
		FindByEmailMock: func(_ context.Context, email string) (*entity.User, error) {
			if email == "test1@dummy.com" {
				return &entity.User{Email: email, PasswordDigest: string(pwDigest)}, nil
			}
			return nil, nil
		},
	}

	tests := []struct {
		name     string
		email    string
		password string
		wantUser entity.User
		wantErr  error
	}{
		{
			name:     "email is not registered",
			email:    "test2@dummy.com",
			password: string(pw),
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.Unauthorized, "email has not been registered"),
		},
		{
			name:     "password is wrong",
			email:    "test1@dummy.com",
			password: "foo",
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.Unauthorized, "failed to authenticate"),
		},
		{
			name:     "email and password is correct",
			email:    "test1@dummy.com",
			password: string(pw),
			wantUser: entity.User{Email: "test1@dummy.com", PasswordDigest: string(pwDigest)},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewLogin(repo)
			user, err := u.Call(context.Background(), tt.email, tt.password)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantUser, user)
		})
	}
}
