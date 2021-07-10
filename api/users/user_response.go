package users

type ResponseUser struct {
	ID         uint   `json:"id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
	AuthToken  string `json:"auth_token"`
}
type ResponseUserLogin struct {
	ID         uint   `json:"id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
	AuthToken  string `json:"auth_token"`
}

type ResponseUserRegistration struct {
	ID                uint   `json:"id"`
	Role              string `json:"role"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ProfilePic        string `json:"profile_pic"`
	RegistrationToken string `json:"registration_token"`
	AuthToken         string `json:"auth_token"`
}

func UserRegistrationResponseFormatter(user UserRegistration, auth_token string) ResponseUserRegistration {
	formatter := ResponseUserRegistration{
		ID:                user.ID,
		Role:              user.Role,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		ProfilePic:        user.ProfilePic,
		RegistrationToken: user.RegistrationToken,
		AuthToken:         auth_token,
	}
	return formatter
}

func UserResponseFormatter(user User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		ID:         user.ID,
		Role:       user.Role,
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		ProfilePic: user.ProfilePic,
		AuthToken:  auth_token,
	}
	return formatter
}
