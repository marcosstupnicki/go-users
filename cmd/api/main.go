package main

import (
	"fmt"
	gologger "github.com/marcosstupnicki/go-logger/pkg"
	"github.com/marcosstupnicki/go-users/cmd/api/handlers"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"os"
)

const (
	ExitCodeFailToCreateWebApplication = iota
	ExitCodeFailToRunWebApplication
)

func main()  {
	app, err := gowebapp.NewWebApplication("local", gologger.DebugLevel)
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


func initRoutes(app *gowebapp.WebApplication){
	userHandler := handlers.NewUserHandler()

	userGroup := app.Group("/users")
	userGroup.Post("/",  userHandler.Create)
	userGroup.Get("/{id}", userHandler.Get)
	userGroup.Put("/{id}", userHandler.Update)
	userGroup.Delete("/{id}", userHandler.Delete)
}

