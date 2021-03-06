package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

type User struct {
	createUsecase *usecase.UserCreate
}

func NewUser(createUsecase *usecase.UserCreate) *User {
	return &User{createUsecase: createUsecase}
}

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userCreateResponse struct {
	User entity.User `json:"user"`
}

// Create godoc
// @Summary Create user
// @ID user-create
// @Accept json
// @Produce json
// @Param payload body userCreateRequest true "User data"
// @Success 201 {object} userCreateResponse
// @Failure 400 {object} errResponse
// @Failure 409 {object} errResponse
// @Router /users [post]
func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var req userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(cerrors.Errorf(cerrors.InvalidInput, err.Error()), w)
		return
	}
	user, err := h.createUsecase.Call(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		handleError(err, w)
		return
	}
	token, err := createToken(user.ID)
	if err != nil {
		handleError(err, w)
		return
	}
	setTokenToCookie(w, token)
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, &userCreateResponse{User: user})
}
