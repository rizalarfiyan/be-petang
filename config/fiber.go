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
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     config.Cors.AllowMethods,
		AllowHeaders:     config.Cors.AllowHeaders,
		AllowCredentials: config.Cors.AllowCredentials,
		ExposeHeaders:    config.Cors.ExposeHeaders,
	}
}
