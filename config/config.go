package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/rizalarfiyan/be-petang/utils"
)

var config *Config

type Config struct {
	Port int
	Host string
	DB   DBConfigs
}

type DBConfigs struct {
	Name               string
	Host               string
	Port               int
	User               string
	Password           string
	ConnectionIdle     time.Duration
	ConnectionLifetime time.Duration
	MaxIdle            int
	MaxOpen            int
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env is not loaded properly: %s", err.Error())
	}

	conf := new(Config)
	conf.Port = utils.GetEnvAsInt("PORT", 8080)
	conf.Host = utils.GetEnv("HOST", "")

	conf.DB.Name = utils.GetEnv("DB_NAME", "")
	conf.DB.Host = utils.GetEnv("DB_HOST", "")
	conf.DB.Port = utils.GetEnvAsInt("DB_PORT", 5432)
	conf.DB.User = utils.GetEnv("DB_USER", "")
	conf.DB.Password = utils.GetEnv("DB_PASSWORD", "")
	conf.DB.ConnectionIdle = utils.GetEnvAsTimeDuration("DB_CONNECTION_IDLE", 1*time.Minute)
	conf.DB.ConnectionLifetime = utils.GetEnvAsTimeDuration("DB_CONNECTION_LIFETIME", 5*time.Minute)
	conf.DB.MaxIdle = utils.GetEnvAsInt("DB_MAX_IDLE", 20)
	conf.DB.MaxOpen = utils.GetEnvAsInt("DB_MAX_OPEN", 50)

	config = conf
}

func Get() *Config {
	return config
}
