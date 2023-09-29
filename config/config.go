package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"video-rest-api"`
	Environment string `envconfig:"ENV" default:"dev"`
	Port        int    `envconfig:"PORT" default:"8002" required:"true"`

	DBHost             string `envconfig:"DB_HOST" default:"localhost"`
	DBPort             string `envconfig:"DB_PORT" default:"3306"`
	DBUserName         string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword         string `envconfig:"DB_PASSWORD" default:""`
	DBDatabaseName     string `envconfig:"DB_DBNAME" default:"go_test"`
	DBLogMode          int    `envconfig:"DB_LOG_MODE" default:"3"`
	GoogleClientId     string `envconfig:"GOGLE_CLIENT_ID"`
	GoogleClientSecret string `envconfig:"GOGLE_CLIENT_ID"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
