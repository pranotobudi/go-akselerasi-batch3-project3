package users

import (
	"log"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/middleware"

	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []common.Route {
	log.Println("INSIDE UserRoutes.Route")
	userHandler := NewUserHandler()

	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: userHandler.UserRegistrationSendEmail,
		},
		{
			Method:  echo.GET,
			Path:    "/register/confirmation",
			Handler: userHandler.RegisterConfirmation,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: userHandler.UserLogin,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id/profile",
			Handler: userHandler.UpdateUserProfile,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user"),
			},
		},

		// CRUD USERS

		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.AddUser,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUser,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUserProfile,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
	}
}
