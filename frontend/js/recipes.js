function fetchRecipes() {
    // Fetch the recipes from the API
    fetch('/recipes')
        .then(response => response.json()) // Parse the JSON from the response
        .then(data => {
            // Get the recipesList element from the HTML
            const recipesList = document.getElementById('recipesList');
            recipesList.innerHTML = ''; // Clear any existing items in the list

            // Iterate over each recipe in the data array
            data.forEach(recipe => {
                // Create a new list item for each recipe
                const listItem = document.createElement('li');
                listItem.textContent = `${recipe.Name} - Rating: ${recipe.RecipeRating}`;
                recipesList.appendChild(listItem); // Add the list item to the list
            });
        })
        .catch(error => console.error('Error fetching recipes:', error)); // Log errors to the console
}

// Load recipes when the page loads
window.onload = fetchRecipes;