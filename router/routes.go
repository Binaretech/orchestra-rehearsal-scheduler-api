package router

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/api"
)

func Routes(router *Router) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		api.ResponseJson(w, 200, map[string]string{
			"message": "hello world",
		})
	})
}
