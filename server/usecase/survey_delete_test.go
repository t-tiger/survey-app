package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func TestSurveyDelete_Call(t *testing.T) {
	repo := &surveyRepoMock{
		FindByMock: func(_ context.Context, id string) (entity.Survey, error) {
			if id == "s1" {
				return entity.Survey{ID: id, PublisherID: "u1"}, nil
			}
			return entity.Survey{}, cerrors.Errorf(cerrors.NotFound, "not found")
		},
		DeleteMock: func(_ context.Context, _ entity.Survey) error {
			return nil
		},
	}

	tests := []struct {
		name    string
		id      string
		userID  string
		wantErr error
	}{
		{
			name:    "id is not found",
			id:      "s2",
			userID:  "u1",
			wantErr: cerrors.Errorf(cerrors.NotFound, "not found"),
		},
		{
			name:    "publisher_id is different",
			id:      "s1",
			userID:  "u2",
			wantErr: cerrors.Errorf(cerrors.Forbidden, "prohibited to delete survey"),
		},
		{
			name:    "succeed to delete",
			id:      "s1",
			userID:  "u1",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSurveyDelete(repo)
			err := u.Call(context.Background(), tt.id, tt.userID)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
