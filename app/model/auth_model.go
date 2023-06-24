package model

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type JWTAuthPayload struct {
	ID       *uuid.UUID `json:"id"`
	Email    string     `json:"email"`
	SureName string     `json:"sure_name"`
	FullName string     `json:"full_name"`
	IsNew    bool       `json:"is_new"`
}

func (jwt *JWTAuthPayload) GetFromFiber(ctx *fiber.Ctx) error {
	userStr, err := json.Marshal(ctx.Locals("user"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(userStr, &jwt)
	if err != nil {
		return err
	}
	return nil
}
