package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jasonDeichert/GoGrocery/internal/service"
)

type RecipeHandler struct {
	Service *service.RecipeService
}

func NewRecipeController(service *service.RecipeService) *RecipeHandler {
	return &RecipeHandler{Service: service}
}

func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.Service.GetAllRecipes()
	if err != nil {
		http.Error(w, "Error getting recipes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(recipes)
}
