package recipes

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func NewRecipeSeeder(db *gorm.DB) {
	CreateRecipeDBTable(db) //create table first, then seeder
	DBSeedRecipe(db)
}
func DBSeedRecipe(db *gorm.DB) error {
	RecipeDataSeed(db)
	return nil
}

func RecipeDataSeed(db *gorm.DB) {
	statement := "INSERT INTO recipes (description, category, country, recipe_pic, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	db.Exec(statement, "description1", "category1", "country1", "http://recipe_pic_url_recipe1.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "description2", "category2", "country2", "http://recipe_pic_url_recipe2.jpg", faker.Timestamp(), faker.Timestamp())
	db.Exec(statement, "description3", "category3", "country3", "http://recipe_pic_url_recipe3.jpg", faker.Timestamp(), faker.Timestamp())
}

func CreateRecipeDBTable(db *gorm.DB) {
	db.AutoMigrate(Recipe{})

	// Create Fresh Recipe Table
	if (db.Migrator().HasTable(&Recipe{})) {
		fmt.Println("Recipe table exist")
		db.Migrator().DropTable(&Recipe{})
	}
	db.Migrator().CreateTable(&Recipe{})

}
