package recipes

import (
	"log"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/middleware"

	"github.com/labstack/echo/v4"
)

type RecipeRoutes struct{}

func (r RecipeRoutes) Route() []common.Route {
	log.Println("INSIDE RecipeRoutes.Route")
	handler := NewRecipeHandler()

	return []common.Route{
		// CRUD RECIPES
		{
			Method:  echo.POST,
			Path:    "/recipes",
			Handler: handler.AddRecipe,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/recipes/:id",
			Handler: handler.GetRecipe,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user", "guest"),
			},
		},
		{
			Method:  echo.PUT,
			Path:    "/recipes/:id",
			Handler: handler.UpdateRecipe,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user"),
			},
		},
		{
			Method:  echo.DELETE,
			Path:    "/recipes/:id",
			Handler: handler.DeleteRecipe,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/recipes/category",
			Handler: handler.GetRecipeByCategory,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user", "guest"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/recipes/country",
			Handler: handler.GetRecipeByCountry,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "user", "guest"),
			},
		},
	}
}
