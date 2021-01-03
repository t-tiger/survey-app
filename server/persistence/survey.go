package persistence

import (
	"context"
	"sort"

	"github.com/t-tiger/survey/server/cerrors"

	"github.com/t-tiger/survey/server/entity"
	"gorm.io/gorm"
)

type Survey struct {
	db *gorm.DB
}

func NewSurvey(db *gorm.DB) *Survey {
	return &Survey{db: db}
}

func (p *Survey) Count(ctx context.Context) (int, error) {
	var c int64
	if err := p.db.WithContext(ctx).Model(&entity.Survey{}).Count(&c).Error; err != nil {
		return 0, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return int(c), nil
}

func (p *Survey) Find(ctx context.Context, limit, offset int) ([]entity.Survey, error) {
	var surveys []entity.Survey
	err := p.db.WithContext(ctx).
		Preload("Questions").Preload("Questions.Options").Preload("Questions.Options.Answers").
		Limit(limit).Offset(offset).
		Order("created_at desc").Find(&surveys).Error
	if err != nil {
		return nil, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}

	// sort entities by sequence
	for _, s := range surveys {
		sort.Slice(s.Questions, func(i, j int) bool {
			return s.Questions[i].Sequence < s.Questions[j].Sequence
		})
		for _, q := range s.Questions {
			sort.Slice(q.Options, func(i, j int) bool {
				return q.Options[i].Sequence < q.Options[j].Sequence
			})
		}
	}
	return surveys, nil
}
