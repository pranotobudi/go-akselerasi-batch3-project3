package users

import (
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
)

type Services interface {
	AddUserRegistrationSendEmail(req RequestUser, regToken string) (UserRegistration, error)
}

type services struct {
	repository  Repository
	taskService common.BackgroundTask
}

func NewServices(repository Repository, taskService common.BackgroundTask) *services {
	return &services{repository, taskService}
}

func (s *services) AddUserRegistrationSendEmail(req RequestUser, regToken string) (UserRegistration, error) {
	entity := UserRegistration{}
	entity.Role = req.Role
	entity.Username = req.Username
	entity.Email = req.Email
	entity.Password = req.Password
	entity.ProfilePic = req.ProfilePic

	//Schedule Email Sending
	email := req.Email
	toEmail := []string{email}
	emailStruct := common.Email{toEmail, req.Role, regToken}
	s.taskService.AddEmailQueue(emailStruct)

	newEntity, err := s.repository.AddUserRegistration(entity)

	return newEntity, err
}
