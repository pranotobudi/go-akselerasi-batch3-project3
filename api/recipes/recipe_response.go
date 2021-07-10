package recipes

type ResponseRecipe struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Country     string `json:"country"`
	RecipePic   string `json:"recipe_pic"`
}

func RecipeResponseFormatter(recipe Recipe) ResponseRecipe {
	formatter := ResponseRecipe{
		ID:          recipe.ID,
		Description: recipe.Description,
		Category:    recipe.Category,
		Country:     recipe.Country,
		RecipePic:   recipe.RecipePic,
	}
	return formatter
}

func RecipesResponseFormatter(recipes []Recipe) []ResponseRecipe {
	var formatters []ResponseRecipe
	for _, recipe := range recipes {
		formatter := ResponseRecipe{
			ID:          recipe.ID,
			Description: recipe.Description,
			Category:    recipe.Category,
			Country:     recipe.Country,
			RecipePic:   recipe.RecipePic,
		}
		formatters = append(formatters, formatter)
	}
	return formatters
}
