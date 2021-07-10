package recipes

type RequestRecipe struct {
	Description string `json:"description"`
	Category    string `json:"category"`
	Country     string `json:"country"`
	RecipePic   string `json:"recipe_pic"`
}
