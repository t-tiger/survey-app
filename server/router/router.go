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
	userCreateUsecase := usecase.NewUserCreate(userRepo)
	userHandler := handler.NewUser(userCreateUsecase)

	r.Post("/users", userHandler.Create)

	return r
}
