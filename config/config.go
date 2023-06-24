package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/rizalarfiyan/be-petang/utils"
)

type Config struct {
	Port  int
	Host  string
	DB    DBConfigs
	Cors  CorsConfigs
	Redis RedisConfigs
	Email EmailConfigs
	FE    FEConfigs
	JWT   JWTConfigs
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

type CorsConfigs struct {
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
}

type RedisConfigs struct {
	Host            string
	Port            int
	User            string
	Password        string
	ExpiredDuration time.Duration
	DialTimeout     time.Duration
}

type EmailConfigs struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type FEConfigs struct {
	BaseUrl           string
	ChangePasswordUrl string
}

type JWTConfigs struct {
	SecretKey string
	Expired   time.Duration
}

var conf *Config

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".env is not loaded properly: %s", err.Error())
	}

	conf = new(Config)
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

	conf.Cors.AllowOrigins = utils.GetEnv("ALLOW_ORIGINS", "")
	conf.Cors.AllowMethods = utils.GetEnv("ALLOW_METHODS", "")
	conf.Cors.AllowHeaders = utils.GetEnv("ALLOW_HEADERS", "")
	conf.Cors.AllowCredentials = utils.GetEnvAsBool("ALLOW_CREDENTIALS", false)
	conf.Cors.ExposeHeaders = utils.GetEnv("EXPOSE_HEADERS", "")

	conf.Redis.Host = utils.GetEnv("REDIS_HOST", "")
	conf.Redis.Port = utils.GetEnvAsInt("REDIS_PORT", 6379)
	conf.Redis.User = utils.GetEnv("REDIS_USER", "")
	conf.Redis.Password = utils.GetEnv("REDIS_PASSWORD", "")
	conf.Redis.ExpiredDuration = utils.GetEnvAsTimeDuration("REDIS_EXPIRED_DURATION", 15*time.Minute)
	conf.Redis.DialTimeout = utils.GetEnvAsTimeDuration("REDIS_DIAL_TIMEOUT", 5*time.Minute)

	conf.Email.Host = utils.GetEnv("EMAIL_HOST", "")
	conf.Email.Port = utils.GetEnvAsInt("EMAIL_PORT", 587)
	conf.Email.User = utils.GetEnv("EMAIL_USER", "")
	conf.Email.Password = utils.GetEnv("EMAIL_PASSWORD", "")
	conf.Email.From = utils.GetEnv("EMAIL_FROM", "")

	conf.FE.BaseUrl = utils.GetEnv("FE_BASE_URL", "")
	conf.FE.ChangePasswordUrl = utils.GetEnv("CHANGE_PASSWORD_URL", "")

	conf.JWT.SecretKey = utils.GetEnv("JWT_SECRET_KEY", "")
	conf.JWT.Expired = utils.GetEnvAsTimeDuration("JWT_EXPIRED", 5*24*time.Hour)

	utils.Success("Config is loaded successfully")
}

func Get() *Config {
	return conf
}
