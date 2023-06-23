package request

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStructWithContext(context.Background(), &req,
		validation.Field(&req.Email, validation.Required, is.EmailFormat),
		validation.Field(&req.Password, validation.Required),
	)
}

type RegisterRequest struct {
	Email                string `json:"email"`
	FullName             string `json:"full_name"`
	SureName             string `json:"sure_name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (req RegisterRequest) Validate() error {
	return validation.ValidateStructWithContext(context.Background(), &req,
		validation.Field(&req.Email, validation.Required, is.EmailFormat),
		validation.Field(&req.FullName, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.SureName, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.PasswordConfirmation, validation.Required.When(req.Password != req.PasswordConfirmation).Error("Password do not match.")),
	)
}
