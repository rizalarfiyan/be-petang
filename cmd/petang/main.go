package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizalarfiyan/be-petang/config"
	"github.com/rizalarfiyan/be-petang/database"
	"github.com/rizalarfiyan/be-petang/utils"
)

func main() {
	conf := config.Get()
	postgres := database.Postgres()
	defer func() {
		err := postgres.Close()
		if err != nil {
			utils.Error("Error closing postgres database: ", err)
		}
	}()

	app := fiber.New(config.FiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New())

	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := app.Listen(baseUrl)
	if err != nil {
		utils.Error("Error app serve: ", err)
	}
}
