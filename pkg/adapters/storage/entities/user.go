package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID	`gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string
	LastName    string
	Email       string 
	Password    string
	PhoneNumber string
	Address     []Address
}
