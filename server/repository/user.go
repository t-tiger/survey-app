package repository

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
)

type User interface {
	FindBy(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, name, email, password string) (entity.User, error)
}
