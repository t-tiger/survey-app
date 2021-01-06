package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type RespondentFetchList struct {
	repo repository.Respondent
}

func NewRespondentFetchList(repo repository.Respondent) *RespondentFetchList {
	return &RespondentFetchList{repo: repo}
}

func (u *RespondentFetchList) Call(ctx context.Context, surveyIDs []string, email, name string) ([]entity.Respondent, error) {
	if len(email) == 0 {
		return nil, cerrors.Errorf(cerrors.InvalidInput, "email must be present")
	}
	if len(name) == 0 {
		return nil, cerrors.Errorf(cerrors.InvalidInput, "name must be present")
	}
	return u.repo.FindBySurveyIDsWithUserInfo(ctx, surveyIDs, email, name)
}
