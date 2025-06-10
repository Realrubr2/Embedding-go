package scrape

import (
	"database/sql"
	"embeddings/chatgpt"
	"embeddings/tmdb"
	"embeddings/turso"
	"errors"
	"log"
	"fmt"
)

// this function takes the scraped titles and searches them in tmdb to get the content info
func ScrapeToContentShow(extractedData []string, provider string) ([]turso.Content, error) {
	contentArr := []turso.Content{}
	// first we need to fetch the genre to make a map of the id's
	if err := tmdb.FetchGenresShow(); err != nil {
		fmt.Println("Error fetching genres:", err)
		return []turso.Content{}, err
	}
	if len(extractedData) == 0 {
		log.Println("No titles extracted.")
		return nil, errors.New("no titles to fetch")
	}
	for _, title := range extractedData {
		content, err := tmdb.FetchShowByTitle(title,provider)
		if err != nil {
			log.Printf("Error fetching title '%s': %v", title, err)
			continue 
		}
		contentArr = append(contentArr, content)
	}

	return contentArr, nil
}

// this function takes the scraped titles and searches them in tmdb to get the content info
func ScrapeToContentMovie(extractedData []string, provider string) ([]turso.Content, error) {
	contentArr := []turso.Content{}
	// first we need to fetch the genre to make a map of the id's
	if err := tmdb.FetchGenresShow(); err != nil {
		fmt.Println("Error fetching genres:", err)
		return []turso.Content{}, err
	}
	if len(extractedData) == 0 {
		log.Println("No titles extracted.")
		return nil, errors.New("no titles to fetch")
	}
	for _, title := range extractedData {
		content, err := tmdb.FetchMovieByTitle(title,provider)
		if err != nil {
			log.Printf("Error fetching title '%s': %v", title, err)
			continue 
		}
		contentArr = append(contentArr, content)
	}

	return contentArr, nil
}

// gets the content and places it into turso
func contentToTurso(db *sql.DB,content []turso.Content) error{
		for _, c := range content {
			turso.CreateContent(db, c);
			vector, err := chatgpt.GenerateEmbeddings(c.Title,c.Description,c.Genres,c.Type);
			if err != nil {
				log.Fatalf("Error generating embedding for:%s,%s", c.Title, err)
			}
			embeddingObj := turso.Embeddings{
				Content_ID: c.ID,
				Vector:     vector,
			}
			turso.CreateEmbeddings(db, embeddingObj)
		}
		return nil
	}
	
