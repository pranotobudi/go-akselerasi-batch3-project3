package users

import (
	"log"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/auth"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/database"

	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []common.Route {
	log.Println("INSIDE UserRoutes.Route")
	db := database.GetDBInstance()
	InitDBTable(db)
	DBSeed(db)
	// db.AutoMigrate(User{}, Role{}, Permission{}, RolePermission{}, movie.Genre{}, movie.Movie{}, movie.GenreMovie{}, movie.MovieReview{})
	repo := NewRepository(db)
	taskService := common.NewBackgroundTask()
	userService := NewServices(repo, taskService)
	authService := auth.NewAuthService()
	taskService.InitEmailSchedulers()
	userHandler := NewHandler(userService, authService)

	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/register/users",
			Handler: userHandler.UserRegistrationSendEmail,
		},
		{
			Method: echo.GET,
			Path:   "/register/confirmation",
			Handler: userHandler.RegisterConfirmation,
		},
		{
			Method: echo.POST,
			Path:   "/login",
			// Handler: userHandler.UserLogin,
		},

		// CRUD NEWS

		{
			Method: echo.POST,
			Path:   "/recipes",
			// Handler: userHandler.AddNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin", "author"),
			// },
		},
		{
			Method: echo.GET,
			Path:   "/recipes/:id",
			// Handler: newsHandler.GetNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWareWithRedirect(),
			// 	middleware.RoleAccessMiddleware("reader"),
			// },
		},
	}
}
