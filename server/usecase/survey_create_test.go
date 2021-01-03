package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func TestSurveyCreate_Call(t *testing.T) {
	repo := &surveyRepoMock{
		CreateMock: func(_ context.Context, s entity.Survey) (entity.Survey, error) {
			return s, nil
		},
	}

	tests := []struct {
		name    string
		s       entity.Survey
		wantS   entity.Survey
		wantErr error
	}{
		{
			name:    "survey title is blank",
			s:       entity.Survey{Title: ""},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "survey title must be present"),
		},
		{
			name:    "question is empty",
			s:       entity.Survey{Title: "s1", Questions: []entity.Question{}},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "at least one question must be contained"),
		},
		{
			name:    "question title is blank",
			s:       entity.Survey{Title: "s1", Questions: []entity.Question{{Title: ""}}},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "question title must be present"),
		},
		{
			name: "option is empty",
			s: entity.Survey{
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{}},
				},
			},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "all questions must have at least one option"),
		},
		{
			name: "option title is blank",
			s: entity.Survey{
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{{Title: ""}}},
				},
			},
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "option title must be present"),
		},
		{
			name: "succeed to create",
			s: entity.Survey{
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{{Title: "o2"}, {Title: "o3"}}},
				},
			},
			wantS: entity.Survey{
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{{Title: "o2"}, {Title: "o3"}}},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSurveyCreate(repo)
			s, err := u.Call(context.Background(), tt.s)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantS, s)
		})
	}
}
