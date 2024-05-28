package ai

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

type Client struct {
	*openai.Client
}

func NewClient() (*Client, error) {
	return &Client{
		Client: openai.NewClient(viper.GetString("open_api_key")),
	}, nil
}

func (c *Client) GenerateRecipe(ctx context.Context, firstIngredient, secondIngredient string) (*RecipeResponse, error) {
	systemPrompt, err := generatePromptFromFile("ai/prompts/systemPrompt.txt", firstIngredient, secondIngredient)
	if err != nil {
		return nil, err
	}

	prompt, err := generatePromptFromFile("ai/prompts/prompt.txt", firstIngredient, secondIngredient)
	if err != nil {
		return nil, err
	}

	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	resp, err := c.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Messages:    dialogue,
		Temperature: 0.7,
	})

	if err != nil || len(resp.Choices) != 1 {
		return nil, fmt.Errorf("error during completion: %v len(choices): %v", err, len(resp.Choices))
	}

	msg := resp.Choices[0].Message

	parsedResponse, err := NewRecipeResponseFromJSON(msg.Content)
	if err != nil {
		return nil, fmt.Errorf("error during response parsing: %v", err)
	}

	return parsedResponse, nil
}

func generatePromptFromFile(filename, firstIngredient, secondIngredient string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening the file: %v", err)
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "{firstIngredient}", firstIngredient)
		line = strings.ReplaceAll(line, "{secondIngredient}", secondIngredient)
		content.WriteString(line)
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading the file: %v", err)
	}

	return content.String(), nil
}
