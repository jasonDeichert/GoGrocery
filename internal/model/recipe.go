package model

import (
	"fmt"
	"time"

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
	Name         string       `json:"recipeName"`
	RecipeRating int          `json:"recipeRating"`
	Tasks        []RecipeTask `gorm:"foreignKey:RecipeID"` // One-to-many relationship
}

type RecipeTask struct {
	gorm.Model
	RecipeID    uint          `gorm:"not null;constraint:OnDelete:CASCADE;"`
	Description string        `json:"description"`
	Time        time.Duration `json:"time"`
}
