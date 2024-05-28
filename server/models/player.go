package models

import "time"

type Player struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
