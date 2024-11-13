package main

import (
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

func main() {
	server := http.NewServeMux()

	r := router.New(server)

	router.Routes(r)

	r.RegisterRoutes()

	r.Listen(":8080")
}
