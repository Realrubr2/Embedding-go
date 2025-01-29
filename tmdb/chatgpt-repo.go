package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// OpenAI API Key (store securely!)
const openAIKey = ""

// GenerateEmbeddings creates a 1536-dimensional vector for a movie
func GenerateEmbeddings(title string, description string, genres string, content_type string) ([]float64, error) {
	// Combine title, genres, and description to form the embedding input
	text := fmt.Sprintf("%s. Genres: %s. %s type: %s", title, genres, description, content_type)

	// OpenAI API endpoint
	url := "https://api.openai.com/v1/embeddings"

	// Prepare JSON request payload
	requestBody, _ := json.Marshal(map[string]interface{}{
		"input": text,
		"model": "text-embedding-ada-002", // OpenAI's embedding model
	})

	// Create HTTP request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+openAIKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	// Extract embedding vector
	if data, ok := response["data"].([]interface{}); ok {
		if embedding, ok := data[0].(map[string]interface{})["embedding"].([]interface{}); ok {
			floatEmbedding := make([]float64, len(embedding))
			for i, v := range embedding {
				floatEmbedding[i] = v.(float64)
			}
			return floatEmbedding, nil
		}
	}

	return nil, fmt.Errorf("failed to extract embeddings")
}