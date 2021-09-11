package main

import (
	"fmt"
	"github.com/marcosstupnicki/go-logger/logger"
	"github.com/marcosstupnicki/go-users/cmd/api/handlers"
	webapplication "github.com/marcosstupnicki/go-webapplication/pkg"
	"os"
)

const (
	ExitCodeFailToCreateWebApplication = iota
	ExitCodeFailToRunWebApplication
)

func main()  {
	app, err := webapplication.NewWebApplication("local", logger.DebugLevel)
	if err != nil {
		app.Logger.Error(fmt.Sprintln("error creating webapplication.", err))
		os.Exit(ExitCodeFailToCreateWebApplication)
	}

	initRoutes(app)

	app.Logger.Info("launching webapplication")
	err = app.Run()
	if err != nil {
		fmt.Print("error booting application", err)
		os.Exit(ExitCodeFailToRunWebApplication)
	}
}


func initRoutes(app *webapplication.WebApplication){
	userHandler := handlers.NewUserHandler()

	userGroup := app.Group("/users")
	userGroup.Post("/",  userHandler.Create)
	userGroup.Get("/{id}", userHandler.Get)
	userGroup.Put("/{id}", userHandler.Update)
	userGroup.Delete("/{id}", userHandler.Delete)
}

