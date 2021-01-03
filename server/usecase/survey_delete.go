package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/repository"
)

type SurveyDelete struct {
	repo repository.Survey
}

func NewSurveyDelete(repo repository.Survey) *SurveyDelete {
	return &SurveyDelete{repo: repo}
}

func (u *SurveyDelete) Call(ctx context.Context, id string, userID string) error {
	s, err := u.repo.FindBy(ctx, id)
	if err != nil {
		return err
	}
	if s.PublisherID != userID {
		return cerrors.Errorf(cerrors.Forbidden, "prohibited to delete survey")
	}
	return u.repo.Delete(ctx, s)
}
