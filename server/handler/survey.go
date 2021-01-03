package handler

import (
	"net/http"

	"github.com/t-tiger/survey/server/entity"

	"github.com/go-chi/render"
	"github.com/gorilla/schema"
	"github.com/t-tiger/survey/server/usecase"
)

var decoder = schema.NewDecoder()

type Survey struct {
	fetchUsecase *usecase.SurveyFetchList
}

func NewSurvey(fetchUsecase *usecase.SurveyFetchList) *Survey {
	return &Survey{
		fetchUsecase: fetchUsecase,
	}
}

type surveyListRequest struct {
	Page  int `json:"page"`
	Count int `json:"count"`
}

type surveyListResponse struct {
	TotalCount int              `json:"total_count"`
	Items      []surveyListItem `json:"items"`
}

type surveyListItem struct {
	ID        string               `json:"id"`
	Title     string               `json:"title"`
	Questions []surveyListQuestion `json:"questions"`
}

type surveyListQuestion struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	Sequence int                `json:"sequence"`
	Options  []surveyListOption `json:"options"`
}

type surveyListOption struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Sequence  int    `json:"sequence"`
	VoteCount int    `json:"vote_count"`
}

func (h *Survey) List(w http.ResponseWriter, r *http.Request) {
	var req surveyListRequest
	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		handleError(err, w)
		return
	}
	total, ss, err := h.fetchUsecase.Call(r.Context(), req.Page, req.Count)
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyListResponse(total, ss)
	render.JSON(w, r, res)
}

type surveyCreateRequest struct {
	// TODO
}

func (h *Survey) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// newSurveyListResponse builds response value from totalCount and entity.Survey
func newSurveyListResponse(totalCnt int, ss []entity.Survey) surveyListResponse {
	res := surveyListResponse{
		TotalCount: totalCnt,
		Items:      make([]surveyListItem, len(ss)),
	}
	for i, s := range ss {
		item := surveyListItem{
			ID:        s.ID,
			Title:     s.Title,
			Questions: make([]surveyListQuestion, len(s.Questions)),
		}
		for j, q := range s.Questions {
			question := surveyListQuestion{
				ID:       q.ID,
				Title:    q.Title,
				Sequence: q.Sequence,
				Options:  make([]surveyListOption, len(q.Options)),
			}
			for k, o := range q.Options {
				question.Options[k] = surveyListOption{
					ID:        o.ID,
					Title:     o.Title,
					Sequence:  o.Sequence,
					VoteCount: len(o.Answers),
				}
			}
			item.Questions[j] = question
		}
		res.Items[i] = item
	}
	return res
}
