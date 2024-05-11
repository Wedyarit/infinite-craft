
# Infinite Craft: Adoptation

In Infinite Craft, the player can craft various elements using previous ones gained. In Infinite Craft, the player starts with four classical elements—Water, Fire, Wind, and Earth—and combines them into new elements by dragging them from the sidebar and placing them on top of each other.



## API Reference


## Client

### `NewClient()`
Creates a new instance of the OpenAI Client.

**Parameters:** None

**Returns:**
- `*Client`: A pointer to the newly created Client instance.
- `error`: An error if encountered during client creation.

### `GenerateRecipe(firstIngredient string, secondIngredient string) -> RecipeResponse`
Generates a recipe based on provided ingredients using the OpenAI GPT-3.5 model.

**Parameters:**
- `firstIngredient` (string): The first ingredient for the recipe.
- `secondIngredient` (string): The second ingredient for the recipe.

**Returns:**
- `*RecipeResponse`: A pointer to the generated RecipeResponse.
- `error`: An error if encountered during recipe generation.

---

## RecipeResponse

### `NewRecipeResponseFromJSON(response string) -> RecipeResponse`
Creates a new RecipeResponse instance from a JSON string.

**Parameters:**
- `response` (string): JSON string representing the response.

**Returns:**
- `*RecipeResponse`: A pointer to the newly created RecipeResponse instance.
- `error`: An error if encountered during response parsing.

---

## API

### `NewAPI(db *pg.DB) -> *API`
Creates a new instance of the API with provided database connection.

**Parameters:**
- `db` (*pg.DB): The database connection.

**Returns:**
- `*API`: A pointer to the newly created API instance.
- `error`: An error if encountered during API creation.

### `Router() -> chi.Mux`
Returns the router for the API endpoints.

**Parameters:** None

**Returns:**
- `chi.Mux`: The chi router instance.

---

## RecipeResource

### `NewRecipeResource(store RecipeStoreProxy) -> *RecipeResource`
Creates a new RecipeResource instance with provided RecipeStoreProxy.

**Parameters:**
- `store` (RecipeStoreProxy): The RecipeStoreProxy instance.

**Returns:**
- `*RecipeResource`: A pointer to the newly created RecipeResource instance.

### `router() -> chi.Mux`
Returns the router for the recipe endpoints.

**Parameters:** None

**Returns:**
- `chi.Mux`: The chi router instance.

---

## RecipeStoreProxy

### `NewRecipeStoreProxy(store RecipeStore, aiClient *ai.Client) -> *RecipeStoreProxy`
Creates a new RecipeStoreProxy instance with provided RecipeStore and AI Client.

**Parameters:**
- `store` (RecipeStore): The RecipeStore instance.
- `aiClient` (*ai.Client): The AI Client instance.

**Returns:**
- `*RecipeStoreProxy`: A pointer to the newly created RecipeStoreProxy instance.

### `GetByIngredients(firstIngredient string, secondIngredient string) -> Recipe`
Retrieves a recipe by its ingredients, either from the database or generates a new one if not found.

**Parameters:**
- `firstIngredient` (string): The first ingredient.
- `secondIngredient` (string): The second ingredient.

**Returns:**
- `*models.Recipe`: A pointer to the retrieved or generated Recipe.
- `error`: An error if encountered during retrieval or generation.

---

This API reference outlines the available methods and their functionalities for interacting with the provided codebase.

## Rest API Reference

#### Get (or generate) recipe by ingredients

#### Endpoint:
```
GET /api/recipe?first_ingredient={firstIngredient}&second_ingredient={secondIngredient}
```

#### Description:
Generates a recipe based on the provided ingredients.

#### Parameters:
- `first_ingredient` (string, required): The first ingredient for the recipe.
- `second_ingredient` (string, required): The second ingredient for the recipe.

#### Response example:

```json
{
    "id": 377,
    "first_ingredient": "Поезд",
    "second_ingredient": "Морковь",
    "result": "Локомотив",
    "emoji": "🚂"
}
```

#### Status Codes:
- `200 OK`: Successfully generated the recipe.
- `400 Bad Request`: Missing or invalid parameters.
- `500 Internal Server Error`: Server encountered an unexpected condition.

---

This REST API reference provides information on how to interact with the `/api/recipe` endpoint to generate a recipe based on the provided ingredients. The response format is also included to illustrate the structure of the response returned by the API.
## Demo

![](https://i.imgur.com/GjG1V0S.png)

## Authors

- [@wedyarit](https://www.github.com/Wedyarit)

