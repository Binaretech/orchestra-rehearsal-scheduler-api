package handler

import (
	"net/http"
	"strconv"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/cache"
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
	cache       cache.Cache
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

	h.cache.Set(strconv.FormatInt(user.ID, 10), token)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]any{"accessToken": token, "user": user})
}

func (h *AuthHandler) Protected(ctx *router.Context) error {
	return ctx.JSON(200, map[string]string{"message": "Protected"})
}

func (h *AuthHandler) Register(r *router.Router) {
	r.Post("/login", h.Login)
}

func (h *AuthHandler) RegisterProtected(r *router.Group) {
	r.Get("/test", h.Protected)
}
