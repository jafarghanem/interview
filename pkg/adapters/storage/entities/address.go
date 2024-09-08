package entities

import (
	"github.com/google/uuid"
)

type Address struct {
	Street  string  
	City    string
	State   string 
	ZipCode string    
	Country string   
	UserID  uuid.UUID 
	User    *User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
