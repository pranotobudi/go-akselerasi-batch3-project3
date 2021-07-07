package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/auth"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"github.com/thanhpk/randstr"
)

type handler struct {
	service     Services
	authService auth.AuthService
}

func NewHandler(service Services, authService auth.AuthService) *handler {
	return &handler{service, authService}
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
	// UserRegistrationResponse
	data := UserRegistrationResponseFormatter(newEntity, auth_token)

	response := common.ResponseFormatter(http.StatusOK, "success", "sending email for user registration successfull", data)

	return c.JSON(http.StatusOK, response)
}
