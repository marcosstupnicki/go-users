package main

import (
	"fmt"
	"github.com/marcosstupnicki/go-users/cmd/api/handlers"
	"github.com/marcosstupnicki/go-users/internal/db"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"

	"os"
)

const (
	ExitCodeFailToCreateWebApplication = iota
	ExitCodeFailToRunWebApplication
	ExitCodeFailConnectToDB
)

func main()  {
	app, err := gowebapp.NewWebApplication("local")
	if err != nil {
		os.Exit(ExitCodeFailToCreateWebApplication)
	}

	db, err := db.InitDb()
	if err != nil {
		os.Exit(ExitCodeFailConnectToDB)
	}

	repository := users.Repository{DB: db}
	service := users.NewService(repository)

	err = initRoutes(app, service)
	if err != nil {
		os.Exit(ExitCodeFailToCreateWebApplication)
	}

	err = app.Run()
	if err != nil {
		fmt.Print("error booting application", err)
		os.Exit(ExitCodeFailToRunWebApplication)
	}
}

func initRoutes(app *gowebapp.WebApplication, service users.Service) error {
	userHandler := handlers.NewHandler(service)

	userGroup := app.Group("/users")
	userGroup.Post("/",  userHandler.Create)
	userGroup.Get("/{id}", userHandler.Get)
	userGroup.Put("/{id}", userHandler.Update)
	userGroup.Delete("/{id}", userHandler.Delete)

	return nil
}

