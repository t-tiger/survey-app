package usecase

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type RespondentCreate struct {
	respondentRepo repository.Respondent
	surveyRepo     repository.Survey
}

func NewRespondentCreate(respondentRepo repository.Respondent, surveyRepo repository.Survey) *RespondentCreate {
	return &RespondentCreate{respondentRepo: respondentRepo, surveyRepo: surveyRepo}
}

func (u *RespondentCreate) Call(ctx context.Context, r entity.Respondent) (created entity.Respondent, err error) {
	s, err := u.surveyRepo.FindBy(ctx, r.SurveyID)
	if err != nil {
		return created, err
	}
	if err := u.validateAnswers(s, r); err != nil {
		return created, err
	}
	cur, err := u.respondentRepo.FindBy(ctx, r.SurveyID, r.Email, r.Name)
	if err != nil {
		return created, err
	}
	if cur != nil {
		return created, cerrors.Errorf(cerrors.InvalidInput, "respondent has already been registered")
	}
	return u.respondentRepo.Create(ctx, r)
}

// validateAnswers confirms whether answers contained in respondent correspond to all questions in the surveys.
func (u *RespondentCreate) validateAnswers(s entity.Survey, r entity.Respondent) error {
	expectedNum := len(s.Questions)
	num := 0

	optionMap := map[string]string{}
	for _, q := range s.Questions {
		for _, o := range q.Options {
			optionMap[o.ID] = q.ID
		}
	}

	answeredQuestions := map[string]bool{}
	for _, a := range r.Answers {
		if qID, ok := optionMap[a.OptionID]; ok {
			num += 1
			if answeredQuestions[qID] {
				return cerrors.Errorf(cerrors.InvalidInput, "multiple answers are contained in a single question")
			}
			answeredQuestions[qID] = true
		}
	}

	if num != expectedNum {
		return cerrors.Errorf(cerrors.InvalidInput, "all questions have to be answered")
	}
	return nil
}
