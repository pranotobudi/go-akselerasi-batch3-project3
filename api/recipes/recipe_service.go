package recipes

import (
	"mime/multipart"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/common"
)

type RecipeServices interface {
	UpdateRecipe(userID uint, description string, category string, country string, recipePicHeader *multipart.FileHeader) (*Recipe, error)
	AddRecipe(description string, category string, country string, recipePicHeader *multipart.FileHeader) (*Recipe, error)
	// AddRecipe(req RequestRecipe) (*Recipe, error)
	GetRecipe(recipeID uint) (*Recipe, error)
	DeleteRecipe(recipeID uint) (*Recipe, error)
	GetRecipeByCategory(category string) ([]Recipe, error)
	GetRecipeByCountry(country string) ([]Recipe, error)
}

type services struct {
	repository RecipeRepository
}

func NewRecipeServices() *services {
	repository := NewRecipeRepository()

	return &services{repository}
}

func (s *services) UpdateRecipe(userID uint, description string, category string, country string, recipePicHeader *multipart.FileHeader) (*Recipe, error) {

	recipePicFile, err := recipePicHeader.Open()
	if err != nil {
		return nil, err
	}
	defer recipePicFile.Close()

	recipePicPath, _ := common.UploadFile(uint(userID), recipePicFile, recipePicHeader.Filename)

	recipe, _ := s.repository.GetRecipe(uint(userID))
	recipe.Description = description
	recipe.Category = category
	recipe.Country = country
	recipe.RecipePic = recipePicPath

	newRecipe, err := s.repository.UpdateRecipe(*recipe)
	if err != nil {
		return nil, err
	}
	return newRecipe, nil

}
func (s *services) AddRecipe(description string, category string, country string, recipePicHeader *multipart.FileHeader) (*Recipe, error) {

	recipePicFile, err := recipePicHeader.Open()
	if err != nil {
		return nil, err
	}
	defer recipePicFile.Close()

	recipePicPath, _ := common.UploadFile(1, recipePicFile, recipePicHeader.Filename)

	entity := Recipe{}
	entity.Description = description
	entity.Category = category
	entity.Country = country
	entity.RecipePic = recipePicPath

	newEntity, err := s.repository.AddRecipe(entity)
	if err != nil {
		return nil, err
	}

	return newEntity, nil
}

// func (s *services) AddRecipe(req RequestRecipe) (*Recipe, error) {
// 	entity := Recipe{}
// 	entity.Description = req.Description
// 	entity.Category = req.Category
// 	entity.Country = req.Country
// 	entity.RecipePic = req.RecipePic

// 	newEntity, err := s.repository.AddRecipe(entity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newEntity, nil
// }
func (s *services) GetRecipe(recipeID uint) (*Recipe, error) {
	newEntity, err := s.repository.GetRecipe(recipeID)
	if err != nil {
		return nil, err
	}

	return newEntity, nil
}

func (s *services) DeleteRecipe(recipeID uint) (*Recipe, error) {
	newEntity, err := s.repository.DeleteRecipe(recipeID)
	if err != nil {
		return nil, err
	}

	return newEntity, err
}

func (s *services) GetRecipeByCategory(category string) ([]Recipe, error) {
	entities, err := s.repository.GetRecipeByCategory(category)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("INSIDE SERVICE GetRecipeByCategory: %+v \n", entities)
	return entities, nil
}

func (s *services) GetRecipeByCountry(country string) ([]Recipe, error) {
	entities, err := s.repository.GetRecipeByCountry(country)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("INSIDE SERVICE GetRecipeByCountry: %+v \n", entities)

	return entities, nil
}
