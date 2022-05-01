package config

import "gorm.io/gorm/logger"

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	LogLevel logger.LogLevel
}

type Config struct {
	Database Database
}

type Configs struct {
	Scope map[string]Config
}
