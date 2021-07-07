package users

import "gorm.io/gorm"

type Repository interface {
	AddUserRegistration(entity UserRegistration) (UserRegistration, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddUserRegistration(entity UserRegistration) (UserRegistration, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}
