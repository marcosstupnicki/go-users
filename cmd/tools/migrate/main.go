package main

import (
	"os"

	"github.com/marcosstupnicki/go-users/internal/platform/config"
	"github.com/marcosstupnicki/go-users/internal/users"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
)

const (
	ExitCodeFailToConnectLocalDB = iota
	ExitCodeFailToMigrateModel
	ExitCodeFailCreateRepository
)

func main() {
	cfg, err := config.GetConfigFromScope(gowebapp.Scope{Environment: "local"})
	repo, err := users.NewMySQL(cfg.Database)
	if err != nil {
		os.Exit(ExitCodeFailCreateRepository)
	}

	err = repo.AutoMigrate()
	if err != nil {
		os.Exit(ExitCodeFailToMigrateModel)
	}
}
