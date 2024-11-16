package middleware

import (
	"net/http"
	"strings"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/cache"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/utils"
)

func Auth(cache cache.Cache) router.Middleware {
	return func(ctx *router.Context, next router.HandlerFunc) error {
		r := ctx.Request()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": errors.UNAUTHORIZED})
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": errors.UNAUTHORIZED})
		}

		tokenString := bearerToken[1]

		token, err := utils.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			return ctx.JSON(http.StatusForbidden, map[string]string{"message": errors.FORBIDDEN})
		}

		if ok, _ := cache.Exists(tokenString); !ok {
			return ctx.JSON(http.StatusForbidden, map[string]string{"message": errors.FORBIDDEN})
		}

		return next(ctx)
	}
}
