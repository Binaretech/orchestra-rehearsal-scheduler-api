package handler

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx *router.Context) error {
	var body LoginRequest

	if err := ctx.Parse(&body); err != nil {
		return err
	}

	user, err := h.authService.GetByEmail(body.Email)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": errors.INVALID_CREDENTIALS})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": errors.INVALID_CREDENTIALS})
	}

	token, err := utils.GenerateToken(user.ID, user.Role)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{"accessToken": token, "user": user})
}

func (h *AuthHandler) Register(r *router.Router) {
	r.Post("/login", h.Login)
}

func (h *AuthHandler) RegisterProtected(group *router.Group) {
}
