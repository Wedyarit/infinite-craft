package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"

	"infinite-craft/ai"
	"infinite-craft/database"
)

type API struct {
	Recipe *RecipeResource
	Player *PlayerResource
}

func NewAPI(db *pg.DB) (*API, error) {
	client, err := ai.NewClient()

	if err != nil {
		return nil, err
	}

	RecipeStore := database.NewRecipeStore(db)
	Recipe := NewRecipeResource(*NewRecipeStoreProxy(
		RecipeStore,
		client,
	))

	PlayerStore := database.NewPlayerStore(db)
	Player := NewPlayerResource(PlayerStore)

	api := &API{
		Recipe: Recipe,
		Player: Player,
	}
	return api, nil
}

func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/recipe", a.Recipe.Router())
	r.Mount("/players", a.Player.Router())

	return r
}
