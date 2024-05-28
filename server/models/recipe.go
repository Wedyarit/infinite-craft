package models

import "time"

type Recipe struct {
	ID               int       `json:"id"`
	FirstIngredient  string    `json:"first_ingredient"`
	SecondIngredient string    `json:"second_ingredient"`
	Result           string    `json:"result"`
	Emoji            string    `json:"emoji"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
