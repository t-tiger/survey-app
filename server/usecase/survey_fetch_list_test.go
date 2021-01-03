package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
)

type surveyRepoMock struct {
	repository.Survey
	CountMock func(ctx context.Context) (int, error)
	FindMock  func(ctx context.Context, limit, offset int) ([]entity.Survey, error)
}

func (r *surveyRepoMock) Count(ctx context.Context) (int, error) {
	return r.CountMock(ctx)
}

func (r *surveyRepoMock) Find(ctx context.Context, limit, offset int) ([]entity.Survey, error) {
	return r.FindMock(ctx, limit, offset)
}

func TestSurveyFetchList_Call(t *testing.T) {
	repo := &surveyRepoMock{
		CountMock: func(ctx context.Context) (int, error) {
			return 5, nil
		},
		FindMock: func(ctx context.Context, limit, offset int) ([]entity.Survey, error) {
			return []entity.Survey{{ID: "s1"}, {ID: "s2"}}, nil
		},
	}

	tests := []struct {
		name      string
		page      int
		count     int
		wantTotal int
		wantSS    []entity.Survey
		wantErr   error
	}{
		{
			name:    "page is less than 1",
			page:    0,
			count:   20,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "page must be greater than or equal to 1"),
		},
		{
			name:    "count is less than 1",
			page:    1,
			count:   0,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "count must be greater than or equal to 1"),
		},
		{
			name:    "count is greater than 100",
			page:    1,
			count:   101,
			wantErr: cerrors.Errorf(cerrors.InvalidInput, "count must be less than or equal to 100"),
		},
		{
			name:      "page is greater than 0 and count is between 1 and 100",
			page:      5,
			count:     20,
			wantTotal: 5,
			wantSS:    []entity.Survey{{ID: "s1"}, {ID: "s2"}},
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewSurveyFetchList(repo)
			total, ss, err := u.Call(context.Background(), tt.page, tt.count)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantTotal, total)
			assert.Equal(t, tt.wantSS, ss)
		})
	}
}
