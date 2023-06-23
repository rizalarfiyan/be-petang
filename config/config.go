package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env is not loaded properly: %s", err.Error())
	}

	conf := new(Config)
	conf.DB_HOST = os.Getenv("DB_HOST")
	conf.DB_PORT = os.Getenv("DB_PORT")
	conf.DB_NAME = os.Getenv("DB_NAME")
	conf.DB_USERNAME = os.Getenv("DB_USERNAME")
	conf.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	config = conf
}

func Get() *Config {
	return config
}
