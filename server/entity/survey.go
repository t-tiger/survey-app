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

type Respondent struct {
	ID       string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SurveyID string
	Email    string
	Name     string
}

type Answer struct {
	RespondentID string
	OptionID     string
}
