package main

import (
	"fmt"
	"github.com/marcosstupnicki/go-users/cmd/api/handlers"
	"github.com/marcosstupnicki/go-users/internal/config"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"

	"os"
)

const (
	ExitCodeFailToCreateWebApplication = iota
	ExitCodeFailToRunWebApplication
	ExitCodeFailReadConfigs
	ExitCodeFailCreateUserService
)

func main()  {
	app := gowebapp.NewWebApp("local")

	cfg, err := config.GetConfigFromEnvironment(app.Scope)
	if err != nil {
		fmt.Print(err)
		os.Exit(ExitCodeFailReadConfigs)
	}

	service, err := users.NewService(cfg.Database)
	if err != nil {
		os.Exit(ExitCodeFailCreateUserService)
	}

	initRoutes(app, service)
	if err != nil {
		os.Exit(ExitCodeFailToCreateWebApplication)
	}

	err = app.Run()
	if err != nil {
		fmt.Print("error booting application", err)
		os.Exit(ExitCodeFailToRunWebApplication)
	}
}

func initRoutes(app *gowebapp.WebApp, service users.Service) {
	userHandler := handlers.NewHandler(service)

	userGroup := app.Group("/users")
	userGroup.Post("/",  userHandler.Create)
	userGroup.Get("/{id}", userHandler.Get)
	userGroup.Put("/{id}", userHandler.Update)
	userGroup.Delete("/{id}", userHandler.Delete)
}

