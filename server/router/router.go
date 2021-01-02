package router

import (
	"net/http"

	"github.com/t-tiger/survey/server/usecase"

	"github.com/t-tiger/survey/server/persistence"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/t-tiger/survey/server/handler"
)

func New(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userRepo := persistence.NewUser(db)
	userUsecase := usecase.NewUser(userRepo)
	userHandler := handler.NewUser(userUsecase)

	r.Post("/users", userHandler.Create)

	return r
}
