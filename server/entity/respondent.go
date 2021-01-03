package entity

import "time"

type Respondent struct {
	ID        string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	SurveyID  string
	Email     string
	Name      string
	Answers   []Answer
	CreatedAt time.Time
}

type Answer struct {
	RespondentID string `gorm:"primary_key"`
	OptionID     string `gorm:"primary_key"`
}
