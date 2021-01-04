package usecase

import (
	"context"
	"strings"
	"testing"

	"github.com/t-tiger/survey/server/repository"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func TestSurveyUpdate_Call(t *testing.T) {
	survey := entity.Survey{
		ID:          "s1",
		PublisherID: "u1",
		Title:       "s1",
		Questions: []entity.Question{
			{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
			{Title: "q2", Options: []entity.Option{{Title: "o2"}, {Title: "o3"}}},
		},
	}

	tests := []struct {
		name    string
		s       entity.Survey
		userID  string
		foundS  entity.Survey
		wantS   entity.Survey
		wantErr error
	}{
		{
			name:    "userID is different from publisherID",
			s:       survey,
			userID:  "u2",
			foundS:  entity.Survey{PublisherID: "u1"},
			wantErr: cerrors.Errorf(cerrors.Forbidden, "you don't have permission to update"),
		},
		{
			name:   "current survey has answers",
			s:      survey,
			userID: "u1",
			foundS: entity.Survey{
				PublisherID: "u1",
				Questions: []entity.Question{
					{Options: []entity.Option{{ID: "o1", Answers: []entity.Answer{{OptionID: "o1"}}}}},
				}},
			wantErr: cerrors.Errorf(cerrors.Forbidden, "unable to update with answers"),
		},
		{
			name:    "survey title is blank",
			s:       entity.Survey{ID: "s1", Title: ""},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of survey title must be between 1 and 100"),
		},
		{
			name:    "length of survey title exceeds 100",
			s:       entity.Survey{ID: "s1", Title: strings.Repeat("a", 101)},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of survey title must be between 1 and 100"),
		},
		{
			name:    "question is empty",
			s:       entity.Survey{ID: "s1", Title: "s1", Questions: []entity.Question{}},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "at least one question must be contained"),
		},
		{
			name:    "question title is blank",
			s:       entity.Survey{ID: "s1", Title: "s1", Questions: []entity.Question{{Title: ""}}},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of question title must be between 1 and 100"),
		},
		{
			name:    "length of question title exceeds 100",
			s:       entity.Survey{ID: "s1", Title: "s1", Questions: []entity.Question{{Title: strings.Repeat("a", 101)}}},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of question title must be between 1 and 100"),
		},
		{
			name: "option is empty",
			s: entity.Survey{
				ID:    "s1",
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{}},
				},
			},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "all questions must have at least one option"),
		},
		{
			name: "option title is blank",
			s: entity.Survey{
				ID:    "s1",
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: "o1"}}},
					{Title: "q2", Options: []entity.Option{{Title: ""}}},
				},
			},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of option title must be between 1 and 100"),
		},
		{
			name: "length of option title exceeds 100",
			s: entity.Survey{
				ID:    "s1",
				Title: "s1",
				Questions: []entity.Question{
					{Title: "q1", Options: []entity.Option{{Title: strings.Repeat("a", 101)}}},
				},
			},
			userID:  "u1",
			foundS:  survey,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "length of option title must be between 1 and 100"),
		},
		{
			name:    "succeed to update",
			s:       survey,
			userID:  "u1",
			foundS:  survey,
			wantS:   survey,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &surveyRepoMock{
				WithTransactionMock: func(
					_ context.Context, repos []repository.Transactional, exec func(repos []repository.Transactional) error,
				) error {
					return exec(repos)
				},
				FindByMock: func(_ context.Context, id string) (entity.Survey, error) {
					return tt.foundS, nil
				},
				DeleteMock: func(_ context.Context, _ entity.Survey) error {
					return nil
				},
				CreateMock: func(_ context.Context, s entity.Survey) (entity.Survey, error) {
					return s, nil
				},
			}

			u := NewSurveyUpdate(repo)
			s, err := u.Call(context.Background(), tt.s, tt.userID)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantS, s)
		})
	}
}
