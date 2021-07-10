package users

type RequestUser struct {
	Role       string `json:"role"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profile_pic"`
}

type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
