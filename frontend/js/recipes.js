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
                listItem.textContent = `${recipe.recipeName} - Rating: ${recipe.recipeRating}`;
                recipesList.appendChild(listItem); // Add the list item to the list
            });
        })
        .catch(error => console.error('Error fetching recipes:', error)); // Log errors to the console
}

// This function will be called when the form is submitted
function handleFormSubmit(event) {
    event.preventDefault(); // Prevent the form from being submitted normally

    // Get the form data
    const formData = new FormData(event.target);

    // Convert the form data to an object
    const data = Object.fromEntries(formData.entries());

    if (data.hasOwnProperty('recipeRating')) {
        data.recipeRating = parseInt(data.recipeRating);
    }

    // Send the data to the server
    fetch('/recipes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        .then(response => response.json())
        .then(data => {
            // The server responded with the new recipe
            console.log('Success:', data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}

// Add the event listener to the form
const form = document.getElementById('recipeForm'); // Replace 'yourFormId' with the actual ID of your form
form.addEventListener('submit', handleFormSubmit);

// Load recipes when the page loads
window.onload = fetchRecipes;

