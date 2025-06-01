package util


import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

// translateToDutch using chatgpt to translate YES I KNOW but i have lots of gpt money
// and try using google translate's api with there shitty cli and bullshit abstractions
// If an error occurs, it returns the original text instead.
func TranslateToDutch(text string) string {
	enviromentvars := LoadEnviroment()
	apiKey := enviromentvars[0]
	if apiKey == "" {
		log.Println("Warning: API key is missing. Returning original text.")
		return text
	}

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4o-mini-2024-07-18",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are a helpful translator that translates English text to Dutch.",
				},
				{
					Role:    "user",
					Content: fmt.Sprintf("Translate this to Dutch: %s", text),
				},
			},
			MaxTokens: 1000,
		},
	)

	if err != nil {
		log.Println("Error during translation:", err)
		return text 
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content
	}

	log.Println("Warning: No translation received. Returning original text.")
	return text 
}