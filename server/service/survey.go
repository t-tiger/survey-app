package service

import (
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func ValidateSurvey(s entity.Survey) error {
	// every resource's title must be present
	// also, survey has at least one question and all questions have at least one option
	if len(s.Title) < 1 || len(s.Title) > 100 {
		return cerrors.Errorf(cerrors.InvalidInput, "length of survey title must be between 1 and 100")
	}
	if len(s.Questions) == 0 {
		return cerrors.Errorf(cerrors.InvalidInput, "at least one question must be contained")
	}
	for _, q := range s.Questions {
		if len(q.Title) < 1 || len(q.Title) > 100 {
			return cerrors.Errorf(cerrors.InvalidInput, "length of question title must be between 1 and 100")
		}
		if len(q.Options) == 0 {
			return cerrors.Errorf(cerrors.InvalidInput, "all questions must have at least one option")
		}
		for _, o := range q.Options {
			if len(o.Title) < 1 || len(o.Title) > 100 {
				return cerrors.Errorf(cerrors.InvalidInput, "length of option title must be between 1 and 100")
			}
		}
	}
	return nil
}
