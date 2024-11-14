package handler

import (
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
