package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rizalarfiyan/be-petang/app/exception"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     conf.Cors.AllowOrigins,
		AllowMethods:     conf.Cors.AllowMethods,
		AllowHeaders:     conf.Cors.AllowHeaders,
		AllowCredentials: conf.Cors.AllowCredentials,
		ExposeHeaders:    conf.Cors.ExposeHeaders,
	}
}
