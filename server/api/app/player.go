package app

import (
	"net/http"

	"infinite-craft/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type PlayerStore interface {
	GetLeaders() ([]*models.Player, error)
	GetPlayerByNickname(nickname string) (*models.Player, error)
	SaveOrUpdate(player *models.Player) error
}

type PlayerResource struct {
	Store PlayerStore
}

func NewPlayerResource(store PlayerStore) *PlayerResource {
	return &PlayerResource{
		Store: store,
	}
}

func (rs *PlayerResource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", rs.getLeaders)
	r.Post("/", rs.createOrUpdatePlayer)
	return r
}

type PlayerRequest struct {
	*models.Player
	ProtectedID int `json:"id"`
}

func (d *PlayerRequest) Bind(r *http.Request) error {
	return nil
}

type PlayerResponse struct {
	*models.Player
}

func newPlayerResponse(p *models.Player) *PlayerResponse {
	return &PlayerResponse{
		Player: p,
	}
}

func (rs *PlayerResource) getLeaders(w http.ResponseWriter, r *http.Request) {
	players, err := rs.Store.GetLeaders()
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	if len(players) == 0 {
		players = []*models.Player{}
	}

	type LeaderResponse struct {
		Place    int    `json:"place"`
		Nickname string `json:"nickname"`
		Score    int    `json:"score"`
	}

	var leaderResponses []*LeaderResponse
	for place, player := range players {
		leaderResponses = append(leaderResponses, &LeaderResponse{
			Place:    place + 1,
			Nickname: player.Nickname,
			Score:    player.Score,
		})
	}

	render.Respond(w, r, leaderResponses)
}

func (rs *PlayerResource) createOrUpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerReq := &PlayerRequest{}
	if err := render.Bind(r, playerReq); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	var updatedPlayer *models.Player

	existingPlayer, err := rs.Store.GetPlayerByNickname(playerReq.Nickname)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	if existingPlayer != nil && existingPlayer.Score < playerReq.Score {
		existingPlayer.Score = playerReq.Score
		if err := rs.Store.SaveOrUpdate(existingPlayer); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		updatedPlayer = existingPlayer
	} else if existingPlayer != nil {
		updatedPlayer = existingPlayer
	} else {
		newPlayer := &models.Player{
			Nickname: playerReq.Nickname,
			Score:    playerReq.Score,
		}
		if err := rs.Store.SaveOrUpdate(newPlayer); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		updatedPlayer = newPlayer
	}

	render.Respond(w, r, newPlayerResponse(updatedPlayer))
}
