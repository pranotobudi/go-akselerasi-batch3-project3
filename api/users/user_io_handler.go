package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/auth"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"github.com/thanhpk/randstr"
)

type handler struct {
	service     UserServices
	authService auth.AuthService
}

func NewUserHandler() *handler {
	userService := NewUserServices()
	authService := auth.NewAuthService()

	return &handler{userService, authService}
}

func (h *handler) UserRegistrationSendEmail(c echo.Context) error {
	// Input Binding
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Add to database
	regToken := randstr.Hex(16) // generate 128-bit hex string
	newEntity, err := h.service.AddUserRegistrationSendEmail(*req, regToken)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(req.Role, newEntity.ID)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}
	// Success UserRegistrationResponse
	data := UserRegistrationResponseFormatter(newEntity, auth_token)

	response := common.ResponseFormatter(http.StatusOK, "success", "sending email for user registration successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) RegisterConfirmation(c echo.Context) error {
	// Input
	role := c.QueryParam("role")
	email := c.QueryParam("email")
	token := c.QueryParam("token")

	// Process Input
	newUser, err := h.service.RegisterConfirmation(role, email, token)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}
	// fmt.Printf("INSIDE Handler RegisterConfirmation: %+v \n", newUser)
	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(newUser.Role, newUser.ID)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Success RegisterConfirmation Response
	data := UserResponseFormatter(*newUser, auth_token)
	response := common.ResponseFormatter(http.StatusOK, "success", "user registration confirmation successfull, user created", data)
	return c.JSON(http.StatusOK, response)

}

func (h *handler) UserLogin(c echo.Context) error {
	// Input Binding
	userLogin := new(RequestUserLogin)
	if err := c.Bind(userLogin); err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Process Input
	authUser, err := h.service.AuthUser(*userLogin)
	fmt.Println("We're IN HERE: USERLOGIN INSIDE: authUser: ", authUser)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(authUser.Role, authUser.ID)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Success UserLoginResponse
	data := UserResponseFormatter(*authUser, auth_token)
	response := common.ResponseFormatter(http.StatusOK, "success", "user login successfull, user created", data)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateUserProfile(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	role := c.FormValue("role")
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	profPicHeader, err := c.FormFile("profile_pic")
	if err != nil {
		return err
	}

	//Process Input
	// Process Input
	newUser, err := h.service.UpdateUserProfile(uint(userID), role, username, email, password, profPicHeader)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Create JWT Token
	auth_token, _ := h.authService.CreateAccessToken(newUser.Role, newUser.ID)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Success UserResponse
	data := UserResponseFormatter(*newUser, auth_token)
	response := common.ResponseFormatter(http.StatusOK, "success", "user successfully updated", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AddUser(c echo.Context) error {
	// Input Binding
	req := new(RequestUser)

	if err := c.Bind(req); err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Add to database
	newEntity, err := h.service.AddUser(*req)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(newEntity.Role, newEntity.ID)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}
	// Success UserResponse
	data := UserResponseFormatter(*newEntity, auth_token)

	response := common.ResponseFormatter(http.StatusOK, "success", "add user successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetUser(c echo.Context) error {
	// Input Binding
	userID, _ := strconv.Atoi(c.Param("id"))

	// Get User from database
	newEntity, err := h.service.GetUser(uint(userID))
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(newEntity.Role, newEntity.ID)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}
	// Success UserResponse
	data := UserResponseFormatter(*newEntity, auth_token)

	response := common.ResponseFormatter(http.StatusOK, "success", "get user successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteUser(c echo.Context) error {
	// Input Binding
	userID, _ := strconv.Atoi(c.Param("id"))

	// Get User from database
	newEntity, err := h.service.DeleteUser(uint(userID))
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.authService.CreateAccessToken(newEntity.Role, newEntity.ID)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}
	// Success UserResponse
	data := UserResponseFormatter(*newEntity, auth_token)

	response := common.ResponseFormatter(http.StatusOK, "success", "delete user successfull", data)

	return c.JSON(http.StatusOK, response)
}
