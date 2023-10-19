package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	DbUser     string `envconfig:"DB_USER"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbHost     string `envconfig:"DB_HOST"`
	DbPort     string `envconfig:"DB_PORT"`
	DbName     string `envconfig:"DB_NAME"`
	AppPort    string `envconfig:"APP_PORT"`
}

var Env EnvConfig

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		log.Fatal(err.Error())
	}
}
