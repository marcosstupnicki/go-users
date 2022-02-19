package main

import (
	"fmt"
	"github.com/marcosstupnicki/go-users/internal/config"
	"github.com/marcosstupnicki/go-users/internal/db"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"os"
)

const (
	ExitCodeFailToConnectLocalDB = iota
	ExitCodeFailToMigrateModel
)

func main() {
	cfg, err := config.GetConfigFromEnvironment(gowebapp.Scope{Environment: "local"})
	db, err := db.Connect(cfg.Database)
	if err != nil {
		fmt.Sprintln("Unable to connect local db")
		os.Exit(ExitCodeFailToConnectLocalDB)
	}

	err = db.AutoMigrate(&users.User{})
	if err != nil {
		os.Exit(ExitCodeFailToMigrateModel)
	}
}
