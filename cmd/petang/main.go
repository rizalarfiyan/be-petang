package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizalarfiyan/be-petang/app"
	"github.com/rizalarfiyan/be-petang/app/handler"
	"github.com/rizalarfiyan/be-petang/config"
	"github.com/rizalarfiyan/be-petang/database"
	"github.com/rizalarfiyan/be-petang/utils"
)

func init() {
	config.Init()
	database.PostgresInit()
}

func main() {
	conf := config.Get()
	postgres := database.PostgresConnection()
	defer func() {
		err := postgres.Close()
		if err != nil {
			utils.Error("Error closing postgres database: ", err)
		}
	}()

	apps := fiber.New(config.FiberConfig())
	apps.Use(recover.New())
	apps.Use(cors.New(config.CorsConfig()))
	apps.Use(logger.New())

	route := app.NewRouter(apps)

	baseHandler := handler.NewBaseHandler()
	authHandler := handler.NewAuthHandler()
	route.BaseRoute(baseHandler)
	route.AuthRoute(authHandler)

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := apps.Listen(baseUrl)
	if err != nil {
		utils.Error("Error app serve: ", err)
	}
}
