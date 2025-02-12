package tmdb

import (
	"database/sql"
	"embeddings/chatgpt"
	"embeddings/turso"
	"embeddings/util"
	"fmt"
	"time"
)

const BaseURL = "https://api.themoviedb.org/3"

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
	Description string   `json:"overview"`
	Provider    string   `json:"provider,omitempty"`
	ImageLink   string   `json:"image_Path"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var genreMap = make(map[int]string)

func Movies(db *sql.DB) {
	start := time.Now()
	// Fetch the genres
	if err := fetchGenresMovie(); err != nil {
		fmt.Println("Error fetching genres:", err)
		return
	}
	// Fetch popular movies
	movies, err := fetchPopularMovies()
	if err != nil {
		fmt.Println("Error fetching movies:", err)
		return
	}
	// Process each movie
	for i := range movies {
		provider, err := fetchMovieProvider(movies[i].ID)
		if err == nil {
			movies[i].Provider = provider
		}

		content := convertMovieToContent(movies[i])
		turso.CreateContent(db, content)
		err = util.AppendJSONToFile("movies", content)
		if err != nil {
			fmt.Println("error in creating json", err)
		}

		embeddingVector, err := chatgpt.GenerateEmbeddings(movies[i].Title, movies[i].Description, content.Genres, "movie")
		if err != nil {
			fmt.Println("Error generating embedding for:", movies[i].Title, err)
		}

		embeddingObj := turso.Embeddings{
			Content_ID: movies[i].ID,
			Vector:     embeddingVector,
		}

		turso.CreateEmbeddings(db, embeddingObj)
	}
	fmt.Println("Execution Time:", time.Since(start))
}

func Shows(db *sql.DB) {
	start := time.Now()
	// Fetch the genre's
	if err := fetchGenresShow(); err != nil {
		fmt.Println("Error fetching genres:", err)
		return
	}

	//Fetches all the popular shows
	shows, err := fetchPopularShows()
	if err != nil {
		fmt.Println("Error fetching movies:", err)
		return
	}

	// here we handle the logic of entering each show into the db
	for i := range shows {
		// Fetch provider
		provider, err := fetchShowProvider(shows[i].ID)
		if err == nil {
			shows[i].Provider = provider
		}
		if provider != "Unknown" && provider != "" {
			err = util.AppendJSONToFile("shows", shows[i])
			if err != nil {
				fmt.Println("error in creating json", err)
			}
		}
		// content := convertMovieToContent(shows[i])
		turso.CreateContent(db, shows[i])


		embeddingVector, err := chatgpt.GenerateEmbeddings(shows[i].Title, shows[i].Description, shows[i].Genres, shows[i].Type)
		if err != nil {
			fmt.Println("Error generating embedding for:", shows[i].Title, err)
		}

		// Create embedding object
		embeddingObj := turso.Embeddings{
			Content_ID: shows[i].ID,
			Vector:     embeddingVector, // Direct float array
		}

		// Insert embedding into DB
		turso.CreateEmbeddings(db, embeddingObj)
	}

	fmt.Println("Execution Time:", time.Since(start))
}
