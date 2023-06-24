package service

import (
	"github.com/rizalarfiyan/be-petang/app/model"
	"github.com/rizalarfiyan/be-petang/app/request"
	"github.com/rizalarfiyan/be-petang/app/response"
)

type AuthService interface {
	Login(req request.LoginRequest) (*response.AuthTokenResponse, error)
	Register(req request.RegisterRequest) error
	ForgotPassword(req request.ForgotPasswordRequest) error
	CheckForgotPassword(token string) (*response.AuthMeResponse, error)
	ChangePassword(req request.ChangePasswordRequest, token string) error
	Me(user model.JWTAuthPayload) (*response.AuthMeResponse, error)
}
