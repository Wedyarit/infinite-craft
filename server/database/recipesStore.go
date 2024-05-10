package database

import (
	"errors"
	"infinite-crafting/models"

	"github.com/go-pg/pg"
)

type RecipeStore struct {
	db *pg.DB
}

func NewRecipeStore(db *pg.DB) *RecipeStore {
	return &RecipeStore{
		db: db,
	}
}

func (s *RecipeStore) GetByResult(result string) (*models.Recipe, error) {
	p := &models.Recipe{}
	err := s.db.Model(p).
		Where("result = ?", result).
		Limit(1).
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (s *RecipeStore) GetByIngredients(firstIngredient string, secondIngredient string) (*models.Recipe, error) {
	p := &models.Recipe{}
	err := s.db.Model(p).
		Where("first_ingredient = ? AND second_ingredient = ?", firstIngredient, secondIngredient).
		Select()

	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			p = &models.Recipe{}
			err = s.db.Model(p).
				Where("first_ingredient = ? AND second_ingredient = ?", secondIngredient, firstIngredient).
				Select()

			if err != nil {
				if errors.Is(err, pg.ErrNoRows) {
					return nil, nil
				}
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (s *RecipeStore) Create(p *models.Recipe) error {
	_, err := s.db.Model(p).Insert()
	return err
}
