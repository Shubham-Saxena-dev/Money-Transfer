package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	EnvConfigs *config
	once       sync.Once
)

type config struct {
	App app
	Db  database
}

type app struct {
	Environment string `env:"APP_ENV" envDefault:"dev"`
	AppName     string `env:"APP_NAME" envDefault:"qonto"`
	AppPort     string `env:"APP_PORT" envDefault:"8080"`
	AppHost     string `env:"APP_HOST" envDefault:"localhost"`
}

type database struct {
	Driver      string `env:"DATABASE_TYPE" envDefault:"mysql"`
	User        string `env:"MYSQL_USER" envDefault:"root"`
	Password    string `env:"MYSQL_PASSWORD" envDefault:"root"`
	Name        string `env:"MYSQL_DATABASE" envDefault:"qonto"`
	Host        string `env:"MYSQL_HOST" envDefault:"mysqlDb"`
	DatabaseUri string `env:"DATABASE_URL" envDefault:"root:root@tcp(mysqldb:3306)/qonto"`
	Port        int    `env:"MYSQL_PORT" envDefault:"3306"`
}

func InitFromFile(configName string) {
	if EnvConfigs == nil {
		once.Do(
			func() {
				EnvConfigs = loadEnvVariables(configName)
			})
	}
}

func loadEnvVariables(fileName string) *config {

	_, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("File: %v not found.", fileName)
	}

	err = godotenv.Load(fileName)

	if err != nil {
		log.Fatalf("unable to load .env file %v: %v", fileName, err)
	}

	config := config{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Failed to parse env variables:%v", err)
	}
	return &config
}
