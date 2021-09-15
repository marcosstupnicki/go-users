package main

import (
	"fmt"
	gologger "github.com/marcosstupnicki/go-logger/pkg"
	"github.com/marcosstupnicki/go-users/cmd/api/handlers"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	err = initRoutes(app)
	if err != nil {
		app.Logger.Error(fmt.Sprintln("error init routes.", err))
		os.Exit(ExitCodeFailToCreateWebApplication)
	}

	app.Logger.Info("launching webapplication")
	err = app.Run()
	if err != nil {
		fmt.Print("error booting application", err)
		os.Exit(ExitCodeFailToRunWebApplication)
	}
}


func initRoutes(app *gowebapp.WebApplication) error {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	dbname := "users"

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Unable to connect local db")
		return err
	}

	repository := users.Repository{DB: db}
	service := users.NewService(repository)

	userHandler := handlers.NewHandler(service)

	userGroup := app.Group("/users")
	userGroup.Post("/",  userHandler.Create)
	userGroup.Get("/{id}", userHandler.Get)
	userGroup.Put("/{id}", userHandler.Update)
	userGroup.Delete("/{id}", userHandler.Delete)

	return nil
}

