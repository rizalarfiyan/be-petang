package repository

import (
	"github.com/rizalarfiyan/be-petang/app/model"

	"github.com/google/uuid"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*model.UserModel, error)
	CheckUserByEmail(email string) (bool, error)
	UpdatePasswordByEmail(email string, password string) error
	CreateUser(payload model.CreateUserModel) (uuid.UUID, error)
}
