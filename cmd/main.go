package main

import (
	"fmt"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/handler"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

func main() {
	cnf := config.LoadConfig(".")

	server := http.NewServeMux()

	r := router.New()

	r.RegisterRoutes(server)

	fmt.Println("Server is running on port", cnf.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cnf.Port), server)
}

func RegisterHandlers(router *router.Router, handlers ...handler.Handler) {
	protected := router.Group("/")
	for _, h := range handlers {
		h.Register(router)
		h.RegisterProtected(protected)
	}
}
