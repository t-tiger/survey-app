package repository

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
)

type Respondent interface {
	FindBy(ctx context.Context, sID, email, name string) (*entity.Respondent, error)
	Create(ctx context.Context, r entity.Respondent) (entity.Respondent, error)
}
