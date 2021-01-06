package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type Survey struct {
	createUsecase    *usecase.SurveyCreate
	deleteUsecase    *usecase.SurveyDelete
	fetchListUsecase *usecase.SurveyFetchList
	findUsecase      *usecase.SurveyFind
	updateUsecase    *usecase.SurveyUpdate
}

func NewSurvey(
	createUsecase *usecase.SurveyCreate,
	deleteUsecase *usecase.SurveyDelete,
	fetchListUsecase *usecase.SurveyFetchList,
	findUsecase *usecase.SurveyFind,
	updateUsecase *usecase.SurveyUpdate,
) *Survey {
	return &Survey{
		createUsecase:    createUsecase,
		deleteUsecase:    deleteUsecase,
		fetchListUsecase: fetchListUsecase,
		findUsecase:      findUsecase,
		updateUsecase:    updateUsecase,
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
	Publisher surveyPublisher    `json:"publisher"`
	Title     string             `json:"title"`
	Questions []questionResponse `json:"questions"`
	CreatedAt time.Time          `json:"created_at"`
}

type surveyPublisher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
	total, ss, err := h.fetchListUsecase.Call(r.Context(), req.Page, req.Count)
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyListResponse(total, ss)
	render.JSON(w, r, res)
}

func (h *Survey) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	s, err := h.findUsecase.Call(r.Context(), id)
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyResponse(s)
	render.JSON(w, r, res)
}

type surveySaveRequest struct {
	Title     string `json:"title"`
	Questions []struct {
		Title   string `json:"title"`
		Options []struct {
			Title string `json:"title"`
		} `json:"options"`
	} `json:"questions"`
}

func (h *Survey) Create(w http.ResponseWriter, r *http.Request) {
	var req surveySaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(cerrors.Errorf(cerrors.InvalidInput, err.Error()), w)
		return
	}
	userID := r.Context().Value(ctxUserID).(string)
	s, err := h.createUsecase.Call(r.Context(), req.toSurvey("", userID))
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyResponse(s)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, res)
}

func (h *Survey) Update(w http.ResponseWriter, r *http.Request) {
	var req surveySaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(cerrors.Errorf(cerrors.InvalidInput, err.Error()), w)
		return
	}
	id := chi.URLParam(r, "id")
	userID := r.Context().Value(ctxUserID).(string)
	s, err := h.updateUsecase.Call(r.Context(), req.toSurvey(id, userID), userID)
	if err != nil {
		handleError(err, w)
		return
	}
	res := newSurveyResponse(s)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, res)
}

func (h *Survey) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := r.Context().Value(ctxUserID).(string)
	if err := h.deleteUsecase.Call(r.Context(), id, userID); err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		ID: s.ID,
		Publisher: surveyPublisher{
			ID:   s.Publisher.ID,
			Name: s.Publisher.Name,
		},
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
func (r *surveySaveRequest) toSurvey(id, userID string) entity.Survey {
	s := entity.Survey{
		ID:          id,
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
