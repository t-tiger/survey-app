package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
)

func TestRespondentFetchList_Call(t *testing.T) {
	rs := []entity.Respondent{
		{ID: "r1", SurveyID: "s1"}, {ID: "r2", SurveyID: "s2"},
	}
	repo := &respondentMockRepo{
		FindBySurveyIDsWithUserInfoMock: func(_ context.Context, _ []string, _, _ string) ([]entity.Respondent, error) {
			return rs, nil
		},
	}

	tests := []struct {
		name      string
		surveyIDs []string
		email     string
		iName     string
		wantRS    []entity.Respondent
		wantErr   error
	}{
		{
			name:      "surveyIDs is empty",
			surveyIDs: []string{},
			email:     "test@dummy.com",
			iName:     "test",
			wantErr:   cerrors.Errorf(cerrors.InvalidInput, "surveyIds must be present"),
		},
		{
			name:      "email is empty",
			surveyIDs: []string{"r1", "r2"},
			email:     "",
			iName:     "test",
			wantErr:   cerrors.Errorf(cerrors.InvalidInput, "email must be present"),
		},
		{
			name:      "name is empty",
			surveyIDs: []string{"r1", "r2"},
			email:     "test@dummy.com",
			iName:     "",
			wantErr:   cerrors.Errorf(cerrors.InvalidInput, "name must be present"),
		},
		{
			name:      "all fields are present",
			surveyIDs: []string{"r1", "r2"},
			email:     "test@dummy.com",
			iName:     "test",
			wantRS:    rs,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewRespondentFetchList(repo)
			rs, err := u.Call(context.Background(), tt.surveyIDs, tt.email, tt.iName)
			if err != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantRS, rs)
		})
	}
}
