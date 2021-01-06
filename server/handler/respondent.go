package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type Respondent struct {
	createUsecase *usecase.RespondentCreate
	listUsecase   *usecase.RespondentFetchList
}

func NewRespondent(
	createUsecase *usecase.RespondentCreate,
	listUsecase *usecase.RespondentFetchList,
) *Respondent {
	return &Respondent{
		createUsecase: createUsecase,
		listUsecase:   listUsecase,
	}
}

type respondentListRequest struct {
	Email     string   `json:"email"`
	Name      string   `json:"name"`
	SurveyIDs []string `json:"surveyIds"`
}

type respondentResponse struct {
	ID       string `json:"id"`
	SurveyID string `json:"survey_id"`
}

// List godoc
// @Summary List of respondent
// @ID respondent-list
// @Produce json
// @Param email query string true "Answered user's email address"
// @Param name query string true "Answered user's name"
// @Param surveyIds query string true "Comma separated ids of survey"
// @Success 200 {array} respondentResponse
// @Failure 400 {object} errResponse
// @Router /respondents [get]
func (h *Respondent) List(w http.ResponseWriter, r *http.Request) {
	// split comma separated surveyIDs
	q := r.URL.Query()
	if v, ok := q["surveyIds"]; ok {
		q["surveyIds"] = strings.Split(v[0], ",")
	}
	var req respondentListRequest
	if err := decoder.Decode(&req, q); err != nil {
		handleError(cerrors.Errorf(cerrors.InvalidInput, err.Error()), w)
		return
	}
	rs, err := h.listUsecase.Call(r.Context(), req.SurveyIDs, req.Email, req.Name)
	if err != nil {
		handleError(err, w)
		return
	}
	// build response value
	res := make([]respondentResponse, len(rs))
	for i, rr := range rs {
		res[i] = newRespondentResponse(rr)
	}
	render.JSON(w, r, res)
}

type respondentCreateRequest struct {
	SurveyID   string   `json:"survey_id"`
	Email      string   `json:"email"`
	Name       string   `json:"name"`
	OptionsIDs []string `json:"option_ids"`
}

// Create godoc
// @Summary Create respondent for survey
// @ID respondent-create
// @Accept json
// @Produce json
// @Param payload body respondentCreateRequest true "Respondent data"
// @Success 201 {object} respondentResponse
// @Failure 400 {object} errResponse
// @Router /respondents [post]
func (h *Respondent) Create(w http.ResponseWriter, r *http.Request) {
	var req respondentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(cerrors.Errorf(cerrors.InvalidInput, err.Error()), w)
		return
	}
	rr, err := h.createUsecase.Call(r.Context(), req.toRespondent())
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, newRespondentResponse(rr))
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

// newRespondentResponse builds respondentResponse from entity.Respondent
func newRespondentResponse(r entity.Respondent) respondentResponse {
	return respondentResponse{
		ID:       r.ID,
		SurveyID: r.SurveyID,
	}
}
