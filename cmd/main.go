package main

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/handler"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

func main() {
	server := http.NewServeMux()

	r := router.New()

	r.RegisterRoutes(server)

	http.ListenAndServe(":8080", server)
}

func RegisterHandlers(router *router.Router, handlers ...handler.Handler) {
	protected := router.Group("/")
	for _, h := range handlers {
		h.Register(router)
		h.RegisterProtected(protected)
	}
}
