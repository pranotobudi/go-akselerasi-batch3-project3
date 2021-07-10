package recipes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/auth"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
)

type handler struct {
	service     RecipeServices
	authService auth.AuthService
}

func NewRecipeHandler() *handler {
	service := NewRecipeServices()
	authService := auth.NewAuthService()

	return &handler{service, authService}
}

func (h *handler) UpdateRecipe(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	description := c.FormValue("description")
	category := c.FormValue("category")
	country := c.FormValue("country")
	recipePicHeader, err := c.FormFile("recipe_pic")
	if err != nil {
		return err
	}

	// Process Input
	newRecipe, err := h.service.UpdateRecipe(uint(userID), description, category, country, recipePicHeader)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeResponse
	data := RecipeResponseFormatter(*newRecipe)
	response := common.ResponseFormatter(http.StatusOK, "success", "recipe successfully updated", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AddRecipe(c echo.Context) error {
	// userID, _ := strconv.Atoi(c.Param("id"))
	description := c.FormValue("description")
	category := c.FormValue("category")
	country := c.FormValue("country")
	recipePicHeader, err := c.FormFile("recipe_pic")
	if err != nil {
		return err
	}
	// Process Input
	newRecipe, err := h.service.AddRecipe(description, category, country, recipePicHeader)
	if err != nil {
		return common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeResponse
	data := RecipeResponseFormatter(*newRecipe)
	response := common.ResponseFormatter(http.StatusOK, "success", "recipe successfully added", data)

	return c.JSON(http.StatusOK, response)
}

// func (h *handler) AddRecipe(c echo.Context) error {
// 	// Input Binding
// 	req := new(RequestRecipe)

// 	if err := c.Bind(req); err != nil {
// 		return common.ResponseErrorFormatter(c, err)
// 	}

// 	// Add to database
// 	newEntity, err := h.service.AddRecipe(*req)
// 	if err != nil {
// 		common.ResponseErrorFormatter(c, err)
// 	}

// 	// Success RecipeResponse
// 	data := RecipeResponseFormatter(*newEntity)

// 	response := common.ResponseFormatter(http.StatusOK, "success", "add recipe successfull", data)

// 	return c.JSON(http.StatusOK, response)
// }

func (h *handler) GetRecipe(c echo.Context) error {
	// Input Binding
	userID, _ := strconv.Atoi(c.Param("id"))

	// Get Recipe from database
	newEntity, err := h.service.GetRecipe(uint(userID))
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeResponse
	data := RecipeResponseFormatter(*newEntity)

	response := common.ResponseFormatter(http.StatusOK, "success", "get recipe successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteRecipe(c echo.Context) error {
	// Input Binding
	userID, _ := strconv.Atoi(c.Param("id"))

	// Delete Recipe from database
	newEntity, err := h.service.DeleteRecipe(uint(userID))
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeResponse
	data := RecipeResponseFormatter(*newEntity)

	response := common.ResponseFormatter(http.StatusOK, "success", "delete recipe successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetRecipeByCategory(c echo.Context) error {
	// Input Binding
	category := c.QueryParam("category")

	// Get Recipe from database
	entities, err := h.service.GetRecipeByCategory(category)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeByCategoryResponse
	data := RecipesResponseFormatter(entities)

	response := common.ResponseFormatter(http.StatusOK, "success", "get recipe by category successfull", data)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetRecipeByCountry(c echo.Context) error {
	// Input Binding
	country := c.QueryParam("country")

	// Get Recipe from database
	entities, err := h.service.GetRecipeByCountry(country)
	if err != nil {
		common.ResponseErrorFormatter(c, err)
	}

	// Success RecipeByCountryResponse
	data := RecipesResponseFormatter(entities)

	response := common.ResponseFormatter(http.StatusOK, "success", "get recipe by country successfull", data)

	return c.JSON(http.StatusOK, response)
}
