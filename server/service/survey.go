package service

import (
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func ValidateSurvey(s entity.Survey) error {
	// every resource's title must be present
	// also, survey has at least one question and all questions have at least one option
	if len(s.Title) == 0 {
		return cerrors.Errorf(cerrors.InvalidInput, "survey title must be present")
	}
	if len(s.Questions) == 0 {
		return cerrors.Errorf(cerrors.InvalidInput, "at least one question must be contained")
	}
	for _, q := range s.Questions {
		if len(q.Title) == 0 {
			return cerrors.Errorf(cerrors.InvalidInput, "question title must be present")
		}
		if len(q.Options) == 0 {
			return cerrors.Errorf(cerrors.InvalidInput, "all questions must have at least one option")
		}
		for _, o := range q.Options {
			if len(o.Title) == 0 {
				return cerrors.Errorf(cerrors.InvalidInput, "option title must be present")
			}
		}
	}
	return nil
}
