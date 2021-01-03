package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/t-tiger/survey/server/handler"
	"github.com/t-tiger/survey/server/persistence"
	"github.com/t-tiger/survey/server/usecase"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userRepo := persistence.NewUser(db)
	userAuthUsecase := usecase.NewUserAuth(userRepo)
	userCreateUsecase := usecase.NewUserCreate(userRepo)
	userHandler := handler.NewUser(userAuthUsecase, userCreateUsecase)

	surveyRepo := persistence.NewSurvey(db)
	surveyCreateUsecase := usecase.NewSurveyCreate(surveyRepo)
	surveyDeleteUsecase := usecase.NewSurveyDelete(surveyRepo)
	surveyFetchListUsecase := usecase.NewSurveyFetchList(surveyRepo)
	surveyUpdateUsecase := usecase.NewSurveyUpdate(surveyRepo)
	surveyHandler := handler.NewSurvey(
		surveyCreateUsecase, surveyDeleteUsecase, surveyFetchListUsecase, surveyUpdateUsecase,
	)

	respondentRepo := persistence.NewRespondent(db)
	respondentCreateUsecase := usecase.NewRespondentCreate(respondentRepo, surveyRepo)
	respondentHandler := handler.NewRespondent(respondentCreateUsecase)

	r.Post("/login", userHandler.Login)
	r.Post("/users", userHandler.Create)
	r.Get("/surveys", surveyHandler.List)
	r.Post("/respondents", respondentHandler.Create)

	// authentication required group
	r.Group(func(r chi.Router) {
		r.Use(handler.AuthUser)
		r.Post("/surveys", surveyHandler.Create)
		r.Put("/surveys/{id}", surveyHandler.Update)
		r.Delete("/surveys/{id}", surveyHandler.Delete)
	})

	return r
}
