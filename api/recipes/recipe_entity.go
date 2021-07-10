package recipes

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	UserID      uint
	Description string
	Category    string
	Country     string
	RecipePic   string
}
