package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type SurveyFetchList struct {
	repo repository.Survey
}

func NewSurveyFetchList(repo repository.Survey) *SurveyFetchList {
	return &SurveyFetchList{repo: repo}
}

func (u *SurveyFetchList) Call(ctx context.Context, page, count int) (total int, ss []entity.Survey, err error) {
	if page < 1 {
		return total, ss, cerrors.Errorf(cerrors.InvalidInput, "page must be greater than or equal to 1")
	}
	if count < 1 {
		return total, ss, cerrors.Errorf(cerrors.InvalidInput, "count must be greater than or equal to 1")
	}
	if count > 100 {
		return total, ss, cerrors.Errorf(cerrors.InvalidInput, "count must be less than or equal to 100")
	}
	if total, err = u.repo.Count(ctx); err != nil {
		return total, ss, err
	}
	offset := (page - 1) * count
	ss, err = u.repo.Find(ctx, count, offset)
	return total, ss, err
}
