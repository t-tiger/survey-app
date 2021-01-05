package persistence

import (
	"context"
	"sort"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/repository"
	"gorm.io/gorm"
)

type Survey struct {
	transaction
}

func NewSurvey(db *gorm.DB) *Survey {
	return &Survey{transaction: transaction{db}}
}

func (p *Survey) StartTransaction(db *gorm.DB) repository.Transactional {
	return NewSurvey(db)
}

func (p *Survey) Count(ctx context.Context) (int, error) {
	var c int64
	if err := p.db.WithContext(ctx).Model(&entity.Survey{}).Count(&c).Error; err != nil {
		return 0, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return int(c), nil
}

func (p *Survey) Find(ctx context.Context, limit, offset int) ([]entity.Survey, error) {
	var ss []entity.Survey
	err := p.preloadedDB().WithContext(ctx).
		Limit(limit).Offset(offset).Order("created_at desc").
		Find(&ss).Error
	if err != nil {
		return nil, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}

	// sort entities by sequence
	for _, s := range ss {
		sort.Slice(s.Questions, func(i, j int) bool {
			return s.Questions[i].Sequence < s.Questions[j].Sequence
		})
		for _, q := range s.Questions {
			sort.Slice(q.Options, func(i, j int) bool {
				return q.Options[i].Sequence < q.Options[j].Sequence
			})
		}
	}
	return ss, nil
}

func (p *Survey) FindBy(ctx context.Context, id string) (entity.Survey, error) {
	var s entity.Survey
	err := p.preloadedDB().WithContext(ctx).
		Where(&entity.Survey{ID: id}).
		First(&s).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Survey{}, cerrors.Errorf(cerrors.NotFound, err.Error())
		}
		return entity.Survey{}, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return s, nil
}

func (p *Survey) Create(ctx context.Context, s entity.Survey) (entity.Survey, error) {
	if err := p.db.WithContext(ctx).Create(&s).Error; err != nil {
		return entity.Survey{}, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return s, nil
}

func (p *Survey) Delete(ctx context.Context, s entity.Survey) error {
	err := p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// delete answers
		if err := tx.Where("option_id in (?)", s.OptionIDs()).Delete(&entity.Answer{}).Error; err != nil {
			return err
		}
		// delete options
		if err := tx.Where("question_id in (?)", s.QuestionIDs()).Delete(&entity.Option{}).Error; err != nil {
			return err
		}
		// delete questions
		if err := tx.Where(&entity.Question{SurveyID: s.ID}).Delete(&entity.Question{}).Error; err != nil {
			return err
		}
		// delete respondents
		if err := tx.Where(&entity.Respondent{SurveyID: s.ID}).Delete(&entity.Respondent{}).Error; err != nil {
			return err
		}
		// delete survey finally
		return tx.Where(&entity.Survey{ID: s.ID}).Delete(&s).Error
	})
	if err != nil {
		return cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return nil
}

func (p *Survey) preloadedDB() *gorm.DB {
	return p.db.Preload("Publisher").Preload("Questions").Preload("Questions.Options").Preload("Questions.Options.Answers")
}
