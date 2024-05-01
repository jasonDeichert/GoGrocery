package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jasonDeichert/GoGrocery/internal/model"
	"github.com/jasonDeichert/GoGrocery/internal/service"
)

type RecipeController struct {
	Service *service.RecipeService
}

func NewRecipeController(service *service.RecipeService) *RecipeController {
	return &RecipeController{Service: service}
}

func (h *RecipeController) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.Service.GetAllRecipes()
	if err != nil {
		http.Error(w, "Error getting recipes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(recipes)
}

func (h *RecipeController) AddRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe model.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, "Error parsing recipe", http.StatusBadRequest)
		return
	}

	err = h.Service.AddRecipe(&recipe)
	if err != nil {
		http.Error(w, "Error adding recipe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}
