package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

type respondentMockRepo struct {
	FindByMock func(ctx context.Context, sID, email, name string) (*entity.Respondent, error)
	CreateMock func(ctx context.Context, r entity.Respondent) (entity.Respondent, error)
}

func (r *respondentMockRepo) FindBy(ctx context.Context, sID, email, name string) (*entity.Respondent, error) {
	return r.FindByMock(ctx, sID, email, name)
}

func (r *respondentMockRepo) Create(ctx context.Context, res entity.Respondent) (entity.Respondent, error) {
	return r.CreateMock(ctx, res)
}

func TestRespondentCreate_Call(t *testing.T) {
	respondentRepo := &respondentMockRepo{
		FindByMock: func(_ context.Context, sID, email, name string) (*entity.Respondent, error) {
			if sID == "s1" && email == "test1@dummy.com" && name == "test1" {
				return &entity.Respondent{ID: "r1"}, nil
			}
			return nil, nil
		},
		CreateMock: func(_ context.Context, r entity.Respondent) (entity.Respondent, error) {
			return r, nil
		},
	}
	surveyRepo := &surveyRepoMock{
		FindByMock: func(_ context.Context, id string) (entity.Survey, error) {
			return entity.Survey{
				ID: id,
				Questions: []entity.Question{
					{ID: "q1", Options: []entity.Option{{ID: "o1"}, {ID: "o2"}}},
					{ID: "q2", Options: []entity.Option{{ID: "o3"}, {ID: "o4"}}},
					{ID: "q3", Options: []entity.Option{{ID: "o5"}}},
				},
			}, nil
		},
	}

	tests := []struct {
		name    string
		r       entity.Respondent
		wantR   entity.Respondent
		wantErr error
	}{
		{
			name: "multiple answers are contained in a single question",
			r: entity.Respondent{
				SurveyID: "s1", Email: "test2@dummy.com", Name: "test2",
				Answers: []entity.Answer{{OptionID: "o1"}, {OptionID: "o2"}, {OptionID: "o3"}, {OptionID: "o5"}},
			},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "multiple answers are contained in a single question"),
		},
		{
			name: "there are questions which are not answered",
			r: entity.Respondent{
				SurveyID: "s1", Email: "test2@dummy.com", Name: "test2",
				Answers: []entity.Answer{{OptionID: "o1"}},
			},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "all questions have to be answered"),
		},
		{
			name: "respondent already exists",
			r: entity.Respondent{
				SurveyID: "s1", Email: "test1@dummy.com", Name: "test1",
				Answers: []entity.Answer{{OptionID: "o1"}, {OptionID: "o3"}, {OptionID: "o5"}},
			},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "respondent has already been registered"),
		},
		{
			name: "succeed to create",
			r: entity.Respondent{
				SurveyID: "s1", Email: "test2@dummy.com", Name: "test2",
				Answers: []entity.Answer{{OptionID: "o1"}, {OptionID: "o3"}, {OptionID: "o5"}},
			},
			wantR: entity.Respondent{
				SurveyID: "s1", Email: "test2@dummy.com", Name: "test2",
				Answers: []entity.Answer{{OptionID: "o1"}, {OptionID: "o3"}, {OptionID: "o5"}},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewRespondentCreate(respondentRepo, surveyRepo)
			r, err := u.Call(context.Background(), tt.r)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantR, r)
		})
	}

}
