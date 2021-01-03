package repository

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
)

type Survey interface {
	Count(ctx context.Context) (int, error)
	Find(ctx context.Context, limit, offset int) ([]entity.Survey, error)
	FindBy(ctx context.Context, id string) (entity.Survey, error)
	Create(ctx context.Context, s entity.Survey) (entity.Survey, error)
	Delete(ctx context.Context, s entity.Survey) error
}
