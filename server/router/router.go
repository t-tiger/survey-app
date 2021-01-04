package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/t-tiger/survey/server/config"
	"github.com/t-tiger/survey/server/handler"
	"github.com/t-tiger/survey/server/persistence"
	"github.com/t-tiger/survey/server/usecase"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: config.Config.AllowedOrigins,
	}))

	userRepo := persistence.NewUser(db)
	loginUsecase := usecase.NewLogin(userRepo)
	authHandler := handler.NewAuth(loginUsecase)

	userCreateUsecase := usecase.NewUserCreate(userRepo)
	userHandler := handler.NewUser(userCreateUsecase)

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
	respondentFetchListUsecase := usecase.NewRespondentFetchList(respondentRepo)
	respondentHandler := handler.NewRespondent(respondentCreateUsecase, respondentFetchListUsecase)

	r.Get("/check_auth", authHandler.CheckAuth)
	r.Post("/login", authHandler.Login)
	r.Post("/users", userHandler.Create)
	r.Get("/surveys", surveyHandler.List)
	r.Get("/respondents", respondentHandler.List)
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
