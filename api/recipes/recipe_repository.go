package recipes

import (
	"fmt"

	"github.com/pranotobudi/go-akselerasi-batch3-project3/database"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	AddRecipe(recipe Recipe) (*Recipe, error)
	GetRecipe(id uint) (*Recipe, error)
	UpdateRecipe(user Recipe) (*Recipe, error)
	DeleteRecipe(id uint) (*Recipe, error)
	GetRecipeByCategory(category string) ([]Recipe, error)
	GetRecipeByCountry(country string) ([]Recipe, error)
}

type repository struct {
	db *gorm.DB
}

func NewRecipeRepository() *repository {
	db := database.GetDBInstance()
	NewRecipeSeeder(db)

	return &repository{db}
}

func (r *repository) AddRecipe(entity Recipe) (*Recipe, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	err = r.db.First(&entity).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("INSIDE REPOSITORY AddEntity: %+v \n", entity)
	return &entity, nil
}

func (r *repository) GetRecipe(id uint) (*Recipe, error) {
	var entity Recipe
	err := r.db.First(&entity, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository) UpdateRecipe(entity Recipe) (*Recipe, error) {
	oldRecipe, err := r.GetRecipe(entity.ID)
	if err != nil {
		return nil, err
	}
	err = r.db.First(&oldRecipe).Error
	oldRecipe.ID = entity.ID
	oldRecipe.Description = entity.Description
	oldRecipe.Category = entity.Category
	oldRecipe.Country = entity.Country
	oldRecipe.RecipePic = entity.RecipePic

	err = r.db.Save(&oldRecipe).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository) DeleteRecipe(id uint) (*Recipe, error) {
	var recipe Recipe
	err := r.db.Find(&recipe, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Delete(&recipe).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (r *repository) GetRecipeByCategory(category string) ([]Recipe, error) {
	var entities []Recipe
	err := r.db.Find(&entities, "category=?", category).Error
	if err != nil {
		return nil, err
	}
	// fmt.Printf("INSIDE REPOSITORY GetRecipeByCategory: %+v \n", entities)
	return entities, nil
}

func (r *repository) GetRecipeByCountry(country string) ([]Recipe, error) {
	var entities []Recipe
	err := r.db.Find(&entities, "country=?", country).Error
	if err != nil {
		return nil, err
	}
	// fmt.Printf("INSIDE REPOSITORY GetRecipeByCountry: %+v \n", entities)
	return entities, nil
}
