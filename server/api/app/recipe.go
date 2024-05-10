package app

import (
	"context"
	"net/http"

	"infinite-crafting/ai"

	"infinite-crafting/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type RecipeStoreProxy struct {
	Store    RecipeStore
	AIClient *ai.Client
}

func NewRecipeStoreProxy(store RecipeStore, aiClient *ai.Client) *RecipeStoreProxy {
	return &RecipeStoreProxy{
		Store:    store,
		AIClient: aiClient,
	}
}

type RecipeStore interface {
	GetByIngredients(firstIngredient string, secondIngredient string) (*models.Recipe, error)
	GetByResult(result string) (*models.Recipe, error)
	Create(p *models.Recipe) error
}

type RecipeResource struct {
	Store RecipeStoreProxy
}

func NewRecipeResource(store RecipeStoreProxy) *RecipeResource {
	return &RecipeResource{
		Store: store,
	}
}

func (rs *RecipeResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", rs.get)
	return r
}

type RecipeRequest struct {
	*models.Recipe
	ProtectedID int `json:"id"`
}

func (d *RecipeRequest) Bind(r *http.Request) error {
	return nil
}

type RecipeResponse struct {
	*models.Recipe
}

func newRecipeResponse(p *models.Recipe) *RecipeResponse {
	return &RecipeResponse{
		Recipe: p,
	}
}

func (p *RecipeStoreProxy) GetByIngredients(firstIngredient string, secondIngredient string) (*models.Recipe, error) {
	recipe, err := p.Store.GetByIngredients(firstIngredient, secondIngredient)
	if err != nil {
		return nil, err
	}

	if recipe != nil {
		return recipe, nil
	}

	ctx := context.Background()
	recipeResponse, err := p.AIClient.GenerateRecipe(ctx, firstIngredient, secondIngredient)
	if err != nil {
		return nil, err
	}

	existingRecipe, err := p.Store.GetByResult(recipeResponse.Result)
	if err != nil {
		return nil, err
	}

	if existingRecipe != nil {
		recipeResponse.Emoji = existingRecipe.Emoji
	}

	newRecipe := &models.Recipe{
		FirstIngredient:  firstIngredient,
		SecondIngredient: secondIngredient,
		Result:           recipeResponse.Result,
		Emoji:            recipeResponse.Emoji,
	}

	if err := p.Store.Create(newRecipe); err != nil {
		return nil, err
	}

	return newRecipe, nil
}

func (rs *RecipeResource) get(w http.ResponseWriter, r *http.Request) {
	firstIngredient := r.URL.Query().Get("first_ingredient")
	secondIngredient := r.URL.Query().Get("second_ingredient")

	recipe, err := rs.Store.GetByIngredients(firstIngredient, secondIngredient)
	if err != nil {
		render.Render(w, r, ErrRender(err))
	}

	render.Respond(w, r, newRecipeResponse(recipe))
}
