package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type Survey struct {
	createUsecase *usecase.SurveyCreate
	fetchUsecase  *usecase.SurveyFetchList
}

func NewSurvey(createUsecase *usecase.SurveyCreate, fetchUsecase *usecase.SurveyFetchList) *Survey {
	return &Survey{
		createUsecase: createUsecase,
		fetchUsecase:  fetchUsecase,
	}
}

type surveyListRequest struct {
	Page  int `json:"page"`
	Count int `json:"count"`
}

type surveyListResponse struct {
	TotalCount int              `json:"total_count"`
	Items      []surveyResponse `json:"items"`
}

type surveyResponse struct {
	ID        string             `json:"id"`
	Title     string             `json:"title"`
	Questions []questionResponse `json:"questions"`
	CreatedAt time.Time          `json:"created_at"`
}

type questionResponse struct {
	ID       string           `json:"id"`
	Title    string           `json:"title"`
	Sequence int              `json:"sequence"`
	Options  []optionResponse `json:"options"`
}

type optionResponse struct {
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
	Title     string `json:"title"`
	Questions []struct {
		Title   string `json:"title"`
		Options []struct {
			Title string `json:"title"`
		} `json:"options"`
	} `json:"questions"`
}

func (h *Survey) Create(w http.ResponseWriter, r *http.Request) {
	var req surveyCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}
	userID := r.Context().Value(ctxUserID).(string)
	s, err := h.createUsecase.Call(r.Context(), req.toSurvey(userID))
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyResponse(s)
	render.JSON(w, r, res)
}

// newSurveyListResponse builds surveyListResponse from totalCount and slice of entity.Survey
func newSurveyListResponse(totalCnt int, ss []entity.Survey) surveyListResponse {
	res := surveyListResponse{
		TotalCount: totalCnt,
		Items:      make([]surveyResponse, len(ss)),
	}
	for i, s := range ss {
		res.Items[i] = newSurveyResponse(s)
	}
	return res
}

// newSurveyListResponse builds surveyResponse from entity.Survey
func newSurveyResponse(s entity.Survey) surveyResponse {
	res := surveyResponse{
		ID:        s.ID,
		Title:     s.Title,
		Questions: make([]questionResponse, len(s.Questions)),
		CreatedAt: s.CreatedAt,
	}
	for j, q := range s.Questions {
		question := questionResponse{
			ID:       q.ID,
			Title:    q.Title,
			Sequence: q.Sequence,
			Options:  make([]optionResponse, len(q.Options)),
		}
		for k, o := range q.Options {
			question.Options[k] = optionResponse{
				ID:        o.ID,
				Title:     o.Title,
				Sequence:  o.Sequence,
				VoteCount: len(o.Answers),
			}
		}
		res.Questions[j] = question
	}
	return res
}

// toSurvey builds entity.Survey from request and authorized userID
func (r *surveyCreateRequest) toSurvey(userID string) entity.Survey {
	s := entity.Survey{
		PublisherID: userID,
		Title:       r.Title,
		Questions:   make([]entity.Question, len(r.Questions)),
	}
	for i, q := range r.Questions {
		qq := entity.Question{
			Sequence: i + 1,
			Title:    q.Title,
			Options:  make([]entity.Option, len(q.Options)),
		}
		for j, o := range q.Options {
			qq.Options[j] = entity.Option{
				Sequence: j + 1,
				Title:    o.Title,
			}
		}
		s.Questions[i] = qq
	}
	return s
}
