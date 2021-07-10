package users

import (
	"time"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/api/recipes"
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

type User struct {
	gorm.Model
	Role       string
	Username   string
	Email      string
	Password   string
	ProfilePic string
	Recipes    []recipes.Recipe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
