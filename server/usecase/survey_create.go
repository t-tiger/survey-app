package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
	"github.com/t-tiger/survey/server/service"
)

type SurveyCreate struct {
	repo repository.Survey
}

func NewSurveyCreate(repo repository.Survey) *SurveyCreate {
	return &SurveyCreate{repo: repo}
}

func (u *SurveyCreate) Call(ctx context.Context, s entity.Survey) (entity.Survey, error) {
	if err := service.ValidateSurvey(s); err != nil {
		return entity.Survey{}, err
	}
	s, err := u.repo.Create(ctx, s)
	if err != nil {
		return entity.Survey{}, err
	}
	return u.repo.FindBy(ctx, s.ID)
}
