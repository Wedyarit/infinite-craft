package database

import (
	"errors"
	"infinite-craft/models"

	"github.com/go-pg/pg"
)

type PlayerStore struct {
	db *pg.DB
}

func NewPlayerStore(db *pg.DB) *PlayerStore {
	return &PlayerStore{
		db: db,
	}
}

func (ps *PlayerStore) GetLeaders() ([]*models.Player, error) {
	var players []*models.Player
	err := ps.db.Model(&players).Order("score DESC").Limit(10).Select()
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (ps *PlayerStore) GetPlayerByNickname(nickname string) (*models.Player, error) {
	player := new(models.Player)
	err := ps.db.Model(player).Where("nickname = ?", nickname).Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return player, nil
}

func (ps *PlayerStore) SaveOrUpdate(player *models.Player) error {
	if player.ID == 0 {
		_, err := ps.db.Model(player).Returning("*").Insert()
		return err
	}
	_, err := ps.db.Model(player).WherePK().Update()
	return err
}
