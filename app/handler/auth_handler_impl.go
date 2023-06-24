package handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rizalarfiyan/be-petang/app/model"
	"github.com/rizalarfiyan/be-petang/app/repository"
	"github.com/rizalarfiyan/be-petang/app/request"
	"github.com/rizalarfiyan/be-petang/app/response"
	"github.com/rizalarfiyan/be-petang/app/service"
	"github.com/rizalarfiyan/be-petang/config"
	"github.com/rizalarfiyan/be-petang/constant"
	"github.com/rizalarfiyan/be-petang/database"
)

type authHandler struct {
	conf    *config.Config
	service service.AuthService
}

func NewAuthHandler(ctx context.Context, conf *config.Config, postgres *sqlx.DB, redis database.RedisInstance) AuthHandler {
	repo := repository.NewAuthRepository(ctx, conf, postgres, redis)
	return &authHandler{
		conf,
		service.NewAuthService(ctx, conf, repo, redis),
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	data, err := h.service.Login(req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success login!",
		Data:    data,
	})
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {
	var req request.RegisterRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	err = h.service.Register(req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Success register!",
	})
}

func (h *authHandler) ForgotPassword(ctx *fiber.Ctx) error {
	var req request.ForgotPasswordRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	err = h.service.ForgotPassword(req)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success send to email!",
	})
}

func (h *authHandler) CheckForgotPassword(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" || len(token) != constant.AuthKeyLength {
		return response.NewErrorMessage(http.StatusBadRequest, constant.ErrorInvalidToken, nil)
	}

	data, err := h.service.CheckForgotPassword(token)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success check token!",
		Data:    data,
	})
}

func (h *authHandler) ChangePassword(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" || len(token) != constant.AuthKeyLength {
		return response.NewErrorMessage(http.StatusBadRequest, constant.ErrorInvalidToken, nil)
	}

	var req request.ChangePasswordRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.NewBindingError(err)
	}

	err = req.Validate()
	if err != nil {
		return response.NewValidationError(err)
	}

	err = h.service.ChangePassword(req, token)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Password has been changed!",
	})
}

func (h *authHandler) Me(ctx *fiber.Ctx) error {
	var user model.JWTAuthPayload
	err := user.GetFromFiber(ctx)
	if err != nil {
		return err
	}

	data, err := h.service.Me(user)
	if err != nil {
		return err
	}

	return ctx.JSON(response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get user data",
		Data:    data,
	})
}
