package handler

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx *router.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Login success"})
}

func (h *AuthHandler) Register(r *router.Router) {
	r.Post("/login", h.Login)
}

func (h *AuthHandler) RegisterProtected(group *router.Group) {
}
