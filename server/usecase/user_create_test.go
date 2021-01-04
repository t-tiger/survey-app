package usecase

import (
	"context"
	"testing"

	"github.com/t-tiger/survey/server/repository"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"golang.org/x/crypto/bcrypt"
)

type userRepoMock struct {
	repository.User
	FindByEmailMock func(ctx context.Context, email string) (*entity.User, error)
	CreateMock      func(ctx context.Context, name, email, password string) (entity.User, error)
}

func (r *userRepoMock) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.FindByEmailMock(ctx, email)
}

func (r *userRepoMock) Create(ctx context.Context, name, email, password string) (entity.User, error) {
	return r.CreateMock(ctx, name, email, password)
}

func TestUserCreate_Call(t *testing.T) {
	type args struct {
		name     string
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantUser entity.User
		wantErr  error
	}{
		{
			name: "name is blank",
			args: args{
				name:     "",
				email:    "test1@dummy.com",
				password: "a1B2@^$_",
			},
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.InvalidInput, "name must not be empty"),
		},
		{
			name: "too short password",
			args: args{
				name:     "test",
				email:    "test1@dummy.com",
				password: "a1B2",
			},
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.InvalidInput, "password length must be greater than or equal to 5"),
		},
		{
			name: "password format is invalid",
			args: args{
				name:     "test",
				email:    "test1@dummy.com",
				password: "a1B2@^$_あ",
			},
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.InvalidInput, "password format is invalid"),
		},
		{
			name: "email format is invalid",
			args: args{
				name:     "test",
				email:    "@dummy.com",
				password: "a1B2@^$_",
			},
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.InvalidInput, "email format is invalid"),
		},
		{
			name: "duplicated email",
			args: args{
				name:     "test",
				email:    "test1@dummy.com",
				password: "a1B2@^$_",
			},
			wantUser: entity.User{},
			wantErr:  cerrors.Errorf(cerrors.Duplicated, "email has already been registered"),
		},
		{
			name: "user is created successfully",
			args: args{
				name:     "test",
				email:    "test@dummy.com",
				password: "a1B2@^$_",
			},
			wantUser: entity.User{Name: "test", Email: "test@dummy.com"},
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock repository
			repo := &userRepoMock{
				FindByEmailMock: func(_ context.Context, email string) (*entity.User, error) {
					if email == "test1@dummy.com" {
						return &entity.User{Email: "test1@dummy.com"}, nil
					}
					return nil, nil
				},
				CreateMock: func(_ context.Context, name, email, password string) (entity.User, error) {
					assert.Equal(t, tt.args.name, name)
					assert.Equal(t, tt.args.email, email)
					// since password is digested, confirm hashed password matches the original password
					assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(password), []byte(tt.args.password)))
					return entity.User{Name: name, Email: email}, nil
				},
			}

			u := NewUserCreate(repo)
			user, err := u.Call(context.Background(), tt.args.name, tt.args.email, tt.args.password)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantUser, user)
		})
	}
}
