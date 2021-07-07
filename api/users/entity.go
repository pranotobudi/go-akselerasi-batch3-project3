package users

import (
	"time"

	"gorm.io/gorm"
)

type UserRegistration struct {
	gorm.Model
	Role              string
	Username          string
	Email             string
	Password          string
	ProfilePic        string
	RegistrationToken string
	TimeCreated       time.Time
}

type Users struct {
	gorm.Model
	Role       string
	Username   string
	Email      string
	Password   string
	ProfilePic string
	Recipes    []Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Recipe struct {
	gorm.Model
	UserID      uint
	Description string
	Category    string
	Country     string
	RecipePic   string
}
