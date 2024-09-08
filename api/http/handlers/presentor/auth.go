package presenter

import (
	"users/internal/user"

	"github.com/google/uuid"
)

type UserRegisterReq struct {
	FirstName string `json:"first_name" validate:"required" example:"yourname"`
	LastName  string `json:"last_name" validate:"required" example:"yourlastname"`
	Email     string `json:"email" validate:"required" example:"abc@gmail.com"`
	Password  string `json:"password" validate:"required" example:"Abc@123"`
	Phone     string `json:"phone" validate:"required" example:"09337307958"`
}
type UserRegisterConcReq struct {
	ID        uuid.UUID `json:"id" validate:"required" example:"e6387992-ae3d-4c33-9b71-12c98b37cb1d"`
	FirstName string    `json:"first_name" validate:"required" example:"yourname"`
	LastName  string    `json:"last_name" validate:"required" example:"yourlastname"`
	Email     string    `json:"email" validate:"required" example:"abc@gmail.com"`
	Password  string    `json:"password" validate:"required" example:"Abc@123"`
	Phone     string    `json:"phone" validate:"required" example:"09337307958"`
}
type UserLoginReq struct {
	Email    string `json:"email" validate:"required" example:"valid_email@folan.com"`
	Password string `json:"password" validate:"required" example:"Abc@123"`
}

func UserRegisterToUserDomain(up *UserRegisterReq) *user.User {
	return &user.User{
		FirstName: up.FirstName,
		LastName:  up.LastName,
		Email:     up.Email,
		Password:  up.Password,
	}
}

func UserRegisterToUserDomainConc(up *UserRegisterConcReq) *user.User {
	return &user.User{
		ID:        up.ID,
		FirstName: up.FirstName,
		LastName:  up.LastName,
		Email:     up.Email,
		Password:  up.Password,
	}
}
