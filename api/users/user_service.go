package users

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	AddUserRegistrationSendEmail(req RequestUser, regToken string) (UserRegistration, error)
	RegisterConfirmation(role string, email string, regToken string) (*User, error)
	AuthUser(req RequestUserLogin) (*User, error)
	UpdateUserProfile(userID uint, role string, username string, email string, password string, profPicHeader *multipart.FileHeader) (*User, error)
	AddUser(req RequestUser) (*User, error)
	GetUser(userID uint) (*User, error)
	DeleteUser(userID uint) (*User, error)
}

type services struct {
	repository  UserRepository
	taskService common.BackgroundTask
}

func NewUserServices() *services {
	repository := NewUserRepository()
	taskService := common.NewBackgroundTask()
	taskService.InitEmailSchedulers()

	return &services{repository, taskService}
}

func (s *services) AddUserRegistrationSendEmail(req RequestUser, regToken string) (UserRegistration, error) {
	entity := UserRegistration{}
	entity.Role = req.Role
	entity.Username = req.Username
	entity.Email = req.Email
	entity.Password = req.Password
	entity.ProfilePic = req.ProfilePic
	entity.RegistrationToken = regToken

	//Schedule Email Sending
	email := req.Email
	toEmail := []string{email}
	emailStruct := common.Email{
		ToEmail:  toEmail,
		Role:     req.Role,
		RegToken: regToken,
	}
	s.taskService.AddEmailQueue(emailStruct)

	newEntity, err := s.repository.AddUserRegistration(entity)

	return newEntity, err
}

func (s *services) RegisterConfirmation(role string, email string, regToken string) (*User, error) {
	userReg, err := s.repository.GetUserRegistration(email)
	if err != nil {
		log.Printf("\n\nINSIDE SERVICE RegisterConfirmation UserReg Not Found \n\n")
		return nil, err
	}
	if userReg.RegistrationToken != regToken {
		return nil, fmt.Errorf("token is not valid")

	}
	fmt.Printf("INSIDE SERVICE RegisterConfirmation UserReg: %+v \n", *userReg)

	// Add to User Table
	user := new(User)
	user.Role = userReg.Role
	user.Username = userReg.Username
	user.Email = userReg.Email
	user.Password = common.GeneratePassword(userReg.Password)
	user.ProfilePic = userReg.ProfilePic

	newUser, err := s.repository.AddUser(*user)
	fmt.Printf("INSIDE SERVICE RegisterConfirmation User: %+v \n", *user)
	if err != nil {
		return nil, err
	}

	return newUser, nil

}

func (s *services) AuthUser(req RequestUserLogin) (*User, error) {
	username := req.Username
	password := req.Password
	fmt.Println("AUTHUSER CALLED, username: ", username, " password: ", password)

	//check Author Table
	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("username is not registered")
	}

	test, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("COMPARES: %s %s \n", user.Password, string(test))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	return user, nil
}

func (s *services) UpdateUserProfile(userID uint, role string, username string, email string, password string, profPicHeader *multipart.FileHeader) (*User, error) {

	profPicfile, err := profPicHeader.Open()
	if err != nil {
		return nil, err
	}
	defer profPicfile.Close()

	profPicPath, _ := common.UploadFile(uint(userID), profPicfile, profPicHeader.Filename)

	user, _ := s.repository.GetUser(uint(userID))
	user.Role = role
	user.Email = email
	user.Password = common.GeneratePassword(password)
	user.Username = username
	user.ProfilePic = profPicPath

	newUser, err := s.repository.UpdateUser(*user)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}

func (s *services) AddUser(req RequestUser) (*User, error) {
	entity := User{}
	entity.Role = req.Role
	entity.Username = req.Username
	entity.Email = req.Email
	entity.Password = common.GeneratePassword(req.Password)
	entity.ProfilePic = req.ProfilePic

	newEntity, err := s.repository.AddUser(entity)
	if err != nil {
		return nil, err
	}

	return newEntity, nil
}
func (s *services) GetUser(userID uint) (*User, error) {
	newEntity, err := s.repository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return newEntity, nil
}

func (s *services) DeleteUser(userID uint) (*User, error) {
	newEntity, err := s.repository.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return newEntity, err
}
