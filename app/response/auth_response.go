package response

import (
	"github.com/google/uuid"
)

type AuthTokenResponse struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"`
}

type AuthMeResponse struct {
	ID       *uuid.UUID `json:"id"`
	Email    string     `json:"email"`
	SureName string     `json:"sure_name"`
	FullName string     `json:"full_name"`
}
