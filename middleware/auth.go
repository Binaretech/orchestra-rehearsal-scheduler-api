package middleware

import (
	"net/http"
	"strings"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
)

func Auth(ctx *router.Context, next router.HandlerFunc) error {
	r := ctx.Request()

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	tokenString := bearerToken[1]

	token, err := utils.ValidateToken(tokenString)

	if err != nil || !token.Valid {
		return ctx.JSON(http.StatusForbidden, map[string]string{"message": "Forbidden"})
	}

	return next(ctx)
}
