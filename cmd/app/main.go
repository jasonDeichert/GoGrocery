package main

import (
	"log"
	"net/http"

	"github.com/jasonDeichert/GoGrocery/internal/controller"

	"github.com/jasonDeichert/GoGrocery/pkg/db"

	"github.com/jasonDeichert/GoGrocery/internal/service"
)

func main() {
	database := db.Init()
	defer db.Close()

	recipeService := service.NewRecipeService(database)
	recipeController := controller.NewRecipeController(recipeService)

	http.HandleFunc("/recipes", recipeController.GetAllRecipes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
