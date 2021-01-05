package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type SurveyFind struct {
	repo repository.Survey
}

func NewSurveyFind(repo repository.Survey) *SurveyFind {
	return &SurveyFind{repo: repo}
}

func (u *SurveyFind) Call(ctx context.Context, id string) (entity.Survey, error) {
	return u.repo.FindBy(ctx, id)
}
