package handler

import (
	"encoding/json"
	"net/http"

	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type Respondent struct {
	createUsecase *usecase.RespondentCreate
}

func NewRespondent(createUsecase *usecase.RespondentCreate) *Respondent {
	return &Respondent{createUsecase: createUsecase}
}

type respondentCreateRequest struct {
	SurveyID   string   `json:"survey_id"`
	Email      string   `json:"email"`
	Name       string   `json:"name"`
	OptionsIDs []string `json:"option_ids"`
}

func (h *Respondent) Create(w http.ResponseWriter, r *http.Request) {
	var req respondentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}
	_, err := h.createUsecase.Call(r.Context(), req.toRespondent())
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// toRespondent builds entity.Respondent from request
func (r *respondentCreateRequest) toRespondent() entity.Respondent {
	res := entity.Respondent{
		SurveyID: r.SurveyID,
		Email:    r.Email,
		Name:     r.Name,
		Answers:  make([]entity.Answer, len(r.OptionsIDs)),
	}
	for i, oID := range r.OptionsIDs {
		res.Answers[i] = entity.Answer{OptionID: oID}
	}
	return res
}
