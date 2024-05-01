package service

import (
	"github.com/jasonDeichert/GoGrocery/internal/model"
	"gorm.io/gorm"
)

type RecipeService struct {
	DB *gorm.DB
}

func NewRecipeService(db *gorm.DB) *RecipeService {
	return &RecipeService{DB: db}
}

func (s *RecipeService) GetAllRecipes() ([]model.Recipe, error) {
	var recipes []model.Recipe
	result := s.DB.Find(&recipes)
	return recipes, result.Error
}

func (s *RecipeService) AddRecipe(recipe *model.Recipe) error {
	result := s.DB.Create(&recipe)
	return result.Error
}
