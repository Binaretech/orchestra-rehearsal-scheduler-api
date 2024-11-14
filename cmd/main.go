package main

import (
	"fmt"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

func main() {
	config.LoadConfig(".")

	server := http.NewServeMux()

	r := router.New(server)

	router.Routes(r)

	r.RegisterRoutes()

	env := config.GetConfig()

	fmt.Println("Server is running on port", env.Port)

	r.Listen(fmt.Sprint(":", env.Port))
}
