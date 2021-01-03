package persistence

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"gorm.io/gorm"
)

type Respondent struct {
	db *gorm.DB
}

func NewRespondent(db *gorm.DB) *Respondent {
	return &Respondent{db: db}
}

func (p *Respondent) FindBy(ctx context.Context, sID, email, name string) (*entity.Respondent, error) {
	var u entity.Respondent
	err := p.db.WithContext(ctx).
		Where(&entity.Respondent{SurveyID: sID, Email: email, Name: name}).
		First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return &u, nil
}

func (p *Respondent) Create(ctx context.Context, r entity.Respondent) (entity.Respondent, error) {
	if err := p.db.WithContext(ctx).Create(&r).Error; err != nil {
		return entity.Respondent{}, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return r, nil
}
