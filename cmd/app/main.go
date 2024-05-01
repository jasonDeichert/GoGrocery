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

	fs := http.FileServer(http.Dir("../../frontend"))
	http.Handle("/", fs)

	recipeRoute(recipeController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func recipeRoute(recipeController *controller.RecipeController) {

	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			recipeController.GetAllRecipes(w, r)
		case "POST":
			recipeController.AddRecipe(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})
}
