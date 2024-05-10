package ai

import (
	"errors"
	"strings"
)

type RecipeResponse struct {
	Result string `json:"result"`
	Emoji  string `json:"emoji"`
}

func NewRecipeResponseFromJSON(response string) (*RecipeResponse, error) {
	response = strings.ReplaceAll(response, "\n", "")

	parts := strings.Split(response, " ")
	if len(parts) < 2 {
		return nil, errors.New("invalid response format")
	}

	result := strings.Join(parts[:len(parts)-1], " ")

	return &RecipeResponse{
		Result: result,
		Emoji:  parts[len(parts)-1],
	}, nil
}
