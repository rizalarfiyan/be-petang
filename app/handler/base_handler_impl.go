package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-petang/app/response"
	"github.com/rizalarfiyan/be-petang/database"
)

type baseHandler struct{}

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}

func (h *baseHandler) Home(ctx *fiber.Ctx) error {
	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success!",
		Data: map[string]interface{}{
			"app_name": "BE Petang",
			"status": map[string]interface{}{
				"postgres": database.PostgresIsConnected(),
			},
		},
	})
}
