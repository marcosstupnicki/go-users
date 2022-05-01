package config

import (
	"errors"

	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"gorm.io/gorm/logger"
)

var _configs = map[string]Config{
	"local": {
		Database: Database{
			User:     "root",
			Password: "root",
			Host:     "127.0.0.1",
			Port:     "3306",
			Name:     "users",
			LogLevel: logger.Info,
		},
	},
}

func GetConfigFromScope(scope gowebapp.Scope) (Config, error) {
	config, found := _configs[scope.Environment]
	if !found {
		return Config{}, errors.New("config not found for indicated scope")
	}

	return config, nil
}
