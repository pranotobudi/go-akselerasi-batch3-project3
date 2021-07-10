package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/api/recipes"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/api/users"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []common.Handler{
		users.UserRoutes{},
		recipes.RecipeRoutes{},
	}

	var routes []common.Route
	for _, handler := range handlers {
		// log.Println("WE'RE HERE routes: ", handler)
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/api")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		}

	}
}
