package users

import (
	"fmt"
	"log"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUserRegistration(entity UserRegistration) (UserRegistration, error)
	GetUserRegistration(email string) (*UserRegistration, error)
	AddUser(user User) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUser(id uint) (*User, error)
	UpdateUser(user User) (*User, error)
	DeleteUser(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository() *repository {
	db := database.GetDBInstance()
	NewUserSeeder(db)

	return &repository{db}
}

func (r *repository) AddUserRegistration(entity UserRegistration) (UserRegistration, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *repository) GetUserRegistration(email string) (*UserRegistration, error) {
	var registration UserRegistration
	err := r.db.First(&registration, "email=?", email).Error
	if err != nil {
		log.Printf("\n\nUserRegistration NOT FOUND IN THE TABLE\n\n")
		return nil, err
	}
	log.Println("UserRegistration FOUND IN THE TABLE")
	return &registration, nil
}

func (r *repository) AddUser(user User) (*User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	err = r.db.First(&user).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("INSIDE REPOSITORY AddUser: %+v \n", user)
	return &user, nil
}

func (r *repository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.First(&user, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUser(id uint) (*User, error) {
	var entity User
	err := r.db.First(&entity, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository) UpdateUser(user User) (*User, error) {
	oldUser, err := r.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	err = r.db.First(&oldUser).Error
	oldUser.Role = user.Role
	oldUser.Email = user.Email
	oldUser.Password = user.Password
	oldUser.Username = user.Username
	oldUser.ProfilePic = user.ProfilePic

	err = r.db.Save(&oldUser).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) DeleteUser(id uint) (*User, error) {
	var user User
	err := r.db.Find(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&user).Error
	// fmt.Printf("================REPOSITORY NEWS: %+v \n\n", news)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
