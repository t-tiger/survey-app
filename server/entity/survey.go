package entity

import "time"

type Survey struct {
	ID          string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	PublisherID string
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Questions   []Question
}

func (s *Survey) QuestionIDs() []string {
	ids := make([]string, len(s.Questions))
	for i, q := range s.Questions {
		ids[i] = q.ID
	}
	return ids
}

func (s *Survey) OptionIDs() []string {
	var ids []string
	for _, q := range s.Questions {
		for _, o := range q.Options {
			ids = append(ids, o.ID)
		}
	}
	return ids
}

func (s *Survey) HasAnswer() bool {
	for _, q := range s.Questions {
		for _, o := range q.Options {
			if len(o.Answers) > 0 {
				return true
			}
		}
	}
	return false
}

type Question struct {
	ID       string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SurveyID string
	Sequence int
	Title    string
	Options  []Option
}

type Option struct {
	ID         string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	QuestionID string
	Sequence   int
	Title      string
	Answers    []Answer
}
