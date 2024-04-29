package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type RecipeRating int

func (r RecipeRating) Validate() error {
	if r < 1 || r > 5 {
		return fmt.Errorf("rating must be between 1 and 5, got %d", r)
	}
	return nil
}

type Recipe struct {
	gorm.Model
	Name         string
	RecipeRating int
}
