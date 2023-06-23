package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-petang/app/handler"
)

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) Router {
	return &router{
		app: app,
	}
}

func (r *router) BaseRoute(handler handler.BaseHandler) {
	r.app.Get("/", handler.Home)
}
