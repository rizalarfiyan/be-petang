package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rizalarfiyan/be-petang/adapter"
	"github.com/rizalarfiyan/be-petang/app/model"
	"github.com/rizalarfiyan/be-petang/app/repository"
	"github.com/rizalarfiyan/be-petang/app/request"
	"github.com/rizalarfiyan/be-petang/app/response"
	"github.com/rizalarfiyan/be-petang/config"
	"github.com/rizalarfiyan/be-petang/constant"
	"github.com/rizalarfiyan/be-petang/database"
	"github.com/rizalarfiyan/be-petang/utils"
)

type authService struct {
	ctx   context.Context
	conf  *config.Config
	repo  repository.AuthRepository
	redis database.RedisInstance
}

func NewAuthService(ctx context.Context, conf *config.Config, repo repository.AuthRepository, redis database.RedisInstance) AuthService {
	return &authService{
		ctx,
		conf,
		repo,
		redis,
	}
}

func (s *authService) createToken(data model.JWTAuthPayload) (string, error) {
	claims := model.TokenJWT{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.conf.JWT.Expired)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.conf.JWT.SecretKey))
}

func (s *authService) Login(req request.LoginRequest) (*response.AuthTokenResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, response.NewErrorMessage(http.StatusBadRequest, constant.ErrorInvalidEmailOrPassword, nil)
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, response.NewErrorMessage(http.StatusBadRequest, constant.ErrorInvalidEmailOrPassword, nil)
	}

	jwtPayload := model.JWTAuthPayload{
		ID:       &user.ID,
		Email:    user.Email,
		SureName: user.SureName,
		FullName: user.FullName,
		IsNew:    false,
	}

	jwtToken, err := s.createToken(jwtPayload)
	if err != nil {
		return nil, err
	}

	return &response.AuthTokenResponse{
		IsNew: jwtPayload.IsNew,
		Token: jwtToken,
	}, nil
}

func (s *authService) Register(req request.RegisterRequest) error {
	foundUser, err := s.repo.CheckUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if foundUser {
		return response.NewErrorMessage(http.StatusBadRequest, constant.ErrorEmailAlreadyRegistered, nil)
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	payload := model.CreateUserModel{
		Email:    req.Email,
		SureName: req.SureName,
		FullName: req.FullName,
		Password: password,
	}

	_, err = s.repo.CreateUser(payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ForgotPassword(req request.ForgotPasswordRequest) error {
	foundUser, err := s.repo.CheckUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if !foundUser {
		return response.NewErrorMessage(http.StatusBadRequest, constant.ErrorInvalidEmail, nil)
	}

	token, _ := gonanoid.New(constant.AuthKeyLength)

	keySearchEmail := fmt.Sprintf("%s%s:*", constant.RedisKeyAuth, req.Email)
	err = s.redis.DelKeysByPatern(keySearchEmail)
	if err != nil {
		return err
	}

	keyAuth := fmt.Sprintf("%s%s:%s", constant.RedisKeyAuth, req.Email, token)
	err = s.redis.Setxc(keyAuth, constant.AuthExpire, req.Email)
	if err != nil {
		return err
	}

	subject := "Forgot Password"
	var data = make(map[string]interface{})
	data["email"] = req.Email
	data["title"] = subject
	data["verificationCode"] = s.conf.FE.BaseUrl + s.conf.FE.ChangePasswordUrl + "?token=" + token

	payload := model.MailPayload{
		From:     s.conf.Email.From,
		To:       req.Email,
		Subject:  subject,
		Template: constant.TemplateChangePassword,
		Data:     data,
	}

	emailConnection := adapter.EmailConnection()
	return NewEmailService(s.conf, emailConnection).SendEmail(payload)
}

func (s *authService) CheckForgotPassword(token string) (*response.AuthMeResponse, error) {
	keyAuth := fmt.Sprintf("%s*:%s", constant.RedisKeyAuth, token)
	search, err := s.redis.Keys(keyAuth)
	if err != nil {
		return nil, err
	}

	if len(search) < 1 {
		return nil, response.NewErrorMessage(http.StatusUnprocessableEntity, constant.ErrorTokenExpired, nil)
	}

	email, err := s.redis.GetString(search[0])
	if err != nil {
		return nil, err
	}

	users, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, response.NewErrorMessage(http.StatusUnprocessableEntity, constant.ErrorUserAccount, nil)
	}

	return &response.AuthMeResponse{
		ID:       &users.ID,
		Email:    users.Email,
		SureName: users.SureName,
		FullName: users.FullName,
	}, nil
}

func (s *authService) ChangePassword(req request.ChangePasswordRequest, token string) error {
	keyAuth := fmt.Sprintf("%s*:%s", constant.RedisKeyAuth, token)
	search, err := s.redis.Keys(keyAuth)
	if err != nil {
		return err
	}

	if len(search) < 1 {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, constant.ErrorTokenExpired, nil)
	}

	email, err := s.redis.GetString(search[0])
	if err != nil {
		return err
	}

	users, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if users == nil {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, constant.ErrorUserAccount, nil)
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	err = s.repo.UpdatePasswordByEmail(email, password)
	if err != nil {
		return err
	}

	err = s.redis.DelKeysByPatern(keyAuth)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Me(user model.JWTAuthPayload) (*response.AuthMeResponse, error) {
	var resp response.AuthMeResponse
	if strings.TrimSpace(user.Email) == "" {
		return &resp, nil
	}

	data, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if data != nil {
		resp.ID = &data.ID
		resp.Email = data.Email
		resp.SureName = data.SureName
		resp.FullName = data.FullName
	}

	return &resp, nil
}
