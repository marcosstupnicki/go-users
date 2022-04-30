package config

import (
	"errors"
	gowebapp "github.com/marcosstupnicki/go-webapp/pkg"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm/logger"
	"testing"
)

func TestGetConfigFromEnvironment(t *testing.T) {
	var tests = []struct {
		name           string
		scope          gowebapp.Scope
		expectedConfig Config
		expectedError  error
	}{
		{
			name: "Ok - GetConfigFromScope for local scope ",
			scope:          gowebapp.Scope{
				Environment: "local",
			},
			expectedConfig: Config{
				Database: Database{
					User: "root",
					Password: "root",
					Host: "127.0.0.1",
					Port: "3306",
					Name: "users",
					LogLevel: logger.Info,
				},
			},
		},
		{
			name: "Error - GetConfigFromScope for unrecognized scope ",
			scope:          gowebapp.Scope{
				Environment: "unrecognized",
			},
			expectedError: errors.New("config not found for indicated scope"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := GetConfigFromScope(tt.scope)
			require.Equal(t, tt.expectedError, err)
			require.Equal(t, tt.expectedConfig, config)
		})
	}
}
