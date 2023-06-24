package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	CheckForgotPassword(ctx *fiber.Ctx) error
	ForgotPassword(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}
