package main

import (
	"fmt"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
)

func main() {
	server := http.NewServeMux()

	r := router.New(server)

	router.Routes(r)

	r.RegisterRoutes()

	fmt.Println("Server is running on port 8080")

	r.Listen(":8080")
}
