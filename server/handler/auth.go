package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/config"
	"github.com/t-tiger/survey/server/entity"
	"github.com/t-tiger/survey/server/usecase"
)

const (
	ctxUserID       = "userID"
	tokenCookieName = "_survey_app_token"
)

type Auth struct {
	userFindUsecase *usecase.UserFind
	loginUsecase    *usecase.Login
}

func NewAuth(userFindUsecase *usecase.UserFind, loginUsecase *usecase.Login) *Auth {
	return &Auth{userFindUsecase: userFindUsecase, loginUsecase: loginUsecase}
}

type checkAuthResponse struct {
	User *entity.User `json:"user"`
}

// CheckAuth godoc
// @Summary Check user's authentication state
// @ID check-auth
// @Accept json
// @Produce json
// @Success 200 {object} checkAuthResponse
// @Failure 404 {object} errResponse
// @Router /check_auth [get]
func (h *Auth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	var user *entity.User
	if userID := retrieveUserID(r); userID != nil {
		u, err := h.userFindUsecase.Call(r.Context(), *userID)
		if err != nil {
			handleError(err, w)
			return
		}
		user = &u
	}
	render.JSON(w, r, &checkAuthResponse{User: user})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	User entity.User `json:"user"`
}

// Login godoc
// @Summary Authenticate User
// @Describing set authentication token in cookie
// @ID login
// @Accept json
// @Produce json
// @Param payload body loginRequest true "Authentication info"
// @Success 200 {object} loginResponse
// @Failure 401 {object} errResponse
// @Router /login [post]
func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(err, w)
		return
	}
	user, err := h.loginUsecase.Call(r.Context(), req.Email, req.Password)
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
	render.JSON(w, r, &loginResponse{User: user})
}

func (h *Auth) Logout(w http.ResponseWriter, _ *http.Request) {
	// remove token cookie
	c := &http.Cookie{Name: tokenCookieName, Value: "", Path: "/", Expires: time.Unix(0, 0), HttpOnly: true}
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusNoContent)
}

// createToken generates json web token
func createToken(userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	s, err := t.SignedString([]byte(config.Config.SecretKey))
	if err != nil {
		return "", cerrors.Errorf(cerrors.Unexpected, err.Error())
	}
	return s, nil
}

// userIDFromToken returns userID by decoding json web token
func userIDFromToken(token string) *string {
	c := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.SecretKey), nil
	})
	if err != nil {
		return nil
	}
	if v, ok := c["user_id"].(string); ok {
		return &v
	}
	return nil
}

// retrieveUserID returns userID by decoding jwt with Authorization value
func retrieveUserID(r *http.Request) *string {
	// first, retrieve from Authorization Header
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 0 {
		tokenVal := strings.Replace(authHeader, "Bearer ", "", 1)
		return userIDFromToken(tokenVal)
	}
	// then, retrieve from cookie
	cookie, err := r.Cookie(tokenCookieName)
	if err != nil {
		return nil
	}
	return userIDFromToken(cookie.Value)
}

func setTokenToCookie(w http.ResponseWriter, token string) {
	exp := time.Now().Add(24 * time.Hour)
	c := &http.Cookie{Name: tokenCookieName, Value: token, Path: "/", Expires: exp, HttpOnly: true}
	http.SetCookie(w, c)
}

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := retrieveUserID(r)
		if userID == nil {
			handleError(cerrors.Errorf(cerrors.Unauthorized, "unauthorized user"), w)
			return
		}
		// set authorized user's id in context
		ctx := context.WithValue(r.Context(), ctxUserID, *userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
