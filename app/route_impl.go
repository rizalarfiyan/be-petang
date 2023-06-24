package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-petang/app/handler"
	"github.com/rizalarfiyan/be-petang/middleware"
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

func (r *router) AuthRoute(handler handler.AuthHandler) {
	auth := r.app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
	auth.Get("/forgot-password", handler.CheckForgotPassword)
	auth.Post("/forgot-password", handler.ForgotPassword)
	auth.Post("/change-password", handler.ChangePassword)

	protected := auth.Group("/", middleware.NewJWTMiddleware(middleware.JWTConfig{}))
	protected.Get("/me", handler.Me)
}
