package exception

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-petang/app/response"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := http.StatusText(code)

	var fiberError *fiber.Error
	var httpError *response.BaseResponse
	var data interface{} = err.Error()
	if errors.As(err, &fiberError) {
		code = fiberError.Code
		message = http.StatusText(code)
		if strings.EqualFold(err.Error(), message) {
			data = nil
		} else {
			data = err.Error()
		}
	}

	if errors.As(err, &httpError) {
		code = httpError.Code
		data = httpError.Data
		if httpError.Message != "" {
			message = httpError.Message
		} else {
			message = http.StatusText(code)
		}
	}

	return ctx.Status(code).JSON(response.BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
