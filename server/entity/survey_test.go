package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurvey_QuestionIDs(t *testing.T) {
	s := &Survey{Questions: []Question{{ID: "q1"}, {ID: "q2"}}}
	actual := s.QuestionIDs()
	assert.Equal(t, []string{"q1", "q2"}, actual)
}

func TestSurvey_OptionIDs(t *testing.T) {
	s := &Survey{Questions: []Question{
		{Options: []Option{{ID: "o1"}, {ID: "o2"}}},
		{Options: []Option{}},
		{Options: []Option{{ID: "o3"}}},
	}}
	actual := s.OptionIDs()
	assert.Equal(t, []string{"o1", "o2", "o3"}, actual)
}
