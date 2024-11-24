package main

import (
	"fmt"
	"net/http"

	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/cache"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/config"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/db"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/errors"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/handler"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/middleware"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/router"
	"github.com/Binaretech/orchestra-rehearsal-scheduler-api/service"
)

const asciiArt = `
  ____  _____   _____ _    _ ______  _____ _______ _____
/  __ \|  __ \ / ____| |  | |  ____|/ ____|__   __|  __  \    /\
| |  | | |__) | |    | |__| | |__  | (___    | |  | |__) |   /  \
| |  | |  _  /| |    |  __  |  __|  \___ \   | |  |  _  /   / /\ \
| |__| | | \ \| |____| |  | | |____ ____) |  | |  | | \ \  / ____ \
 \____/|_|  \_\\_____|_|  |_|______|_____/   |_|  |_|  \_\/_/    \_\

   _____      __             __      __         ___    ____  ____
  / ___/_____/ /_  ___  ____/ /_  __/ /__  ____/   |  / __ \/  _/
  \__ \/ ___/ __ \/ _ \/ __  / / / / / _ \/ __/ /| | / /_/ // /
 ___/ / /__/ / / /  __/ /_/ / /_/ / /  __/ / / ___ |/ ____// /
/____/\___/_/ /_/\___/\__,_/\__,_/_/\___/_/ /_/  |_/_/   /___/

================================================================`

func main() {
	cnf := config.LoadConfig(".")

	server := http.NewServeMux()

	r := router.New()

	r.SetErrorHandler(errors.Handler)

	db, err := db.Connect()

	if err != nil {
		fmt.Println(err)
		return
	}

	cache := cache.NewFileCache()
	defer cache.Close()

	authService := service.NewAuthService(db)
	sectionService := service.NewSectionService(db)
	instrumentService := service.NewInstrumentService(db)

	authHandler := handler.NewAuthHandler(authService, cache)
	sectionHandler := handler.NewSectionHandler(sectionService)
	instrumentHander := handler.NewInstrumentHandler(instrumentService, sectionService)

	RegisterHandlers(r, cache,
		authHandler,
		sectionHandler,
		instrumentHander,
	)

	r.RegisterRoutes(server)

	fmt.Println(asciiArt)
	fmt.Println("Server is running on port", cnf.Port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", cnf.Port), server)

	if err != nil {
		fmt.Println(err)
	}
}

func RegisterHandlers(router *router.Router, cache cache.Cache, handlers ...handler.Handler) {
	protected := router.Group("", middleware.Auth(cache))

	for _, h := range handlers {
		h.Register(router)
		h.RegisterProtected(protected)
	}
}
