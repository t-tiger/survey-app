package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
	"github.com/t-tiger/survey/server/service"
)

type SurveyUpdate struct {
	repo repository.Survey
}

func NewSurveyUpdate(repo repository.Survey) *SurveyUpdate {
	return &SurveyUpdate{repo: repo}
}

func (u *SurveyUpdate) Call(ctx context.Context, s entity.Survey, userID string) (updated entity.Survey, err error) {
	cur, err := u.repo.FindBy(ctx, s.ID)
	if err != nil {
		return updated, err
	}
	if cur.PublisherID != userID {
		return updated, cerrors.Errorf(cerrors.Forbidden, "you don't have permission to update")
	}
	if cur.HasAnswer() {
		return updated, cerrors.Errorf(cerrors.Forbidden, "unable to update with answers")
	}
	if err := service.ValidateSurvey(s); err != nil {
		return updated, err
	}
	err = u.repo.WithTransaction(ctx, []repository.Transactional{u.repo}, func(repos []repository.Transactional) error {
		repo := repos[0].(repository.Survey)
		if err := repo.Delete(ctx, cur); err != nil {
			return err
		}
		updated, err = repo.Create(ctx, s)
		return err
	})
	if err != nil {
		return updated, nil
	}
	return u.repo.FindBy(ctx, updated.ID)
}
