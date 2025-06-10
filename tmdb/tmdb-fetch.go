package tmdb

import (
	"embeddings/turso"
	"embeddings/util"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func FetchShowByTitle(movieTitle string, provider string) (turso.Content, error) {
	encodedTitle := strings.ReplaceAll(movieTitle, " ", "%20")
	url := fmt.Sprintf("%s/search/show?query=%s&include_adult=true&language=en-US&page=1", BaseURL, encodedTitle)
	data, err := makeRequest(url)
	if err != nil {
		return turso.Content{}, errors.New("fetching tmdb err")
	}
	var response struct {
		Results []struct {
			ID          int    `json:"id"`
			Title       string `json:"name"`
			GenreIDs    []int  `json:"genre_ids"`
			ReleaseDate string `json:"first_air_date"`
			Overview    string `json:"overview"`
			Poster string 		`json:"poster_path"`
			} `json:"results"`
		}
		if err := json.Unmarshal(data, &response); err != nil {
			return turso.Content{}, err
		}

		if len(response.Results) == 0 {
			return turso.Content{}, errors.New("no data found for this movie")
		}
		show := response.Results[0]
		content := turso.Content{
			ID:          show.ID,
			Title:       show.Title,
			Genres:      mapGenreIDToString(show.GenreIDs),
			ReleaseDate: show.ReleaseDate,
		Description: util.TranslateToDutch(show.Overview),
		Provider: provider,
		ImageLink:   show.Poster,
		Type:        "movie",
	}
	
	return content, nil
}



func FetchMovieByTitle(movieTitle string, provider string) (turso.Content, error) {
	encodedTitle := strings.ReplaceAll(movieTitle, " ", "%20")
	url := fmt.Sprintf("%s/search/movie?query=%s&include_adult=true&language=en-US&page=1", BaseURL, encodedTitle)
	data, err := makeRequest(url)
	if err != nil {
		return turso.Content{}, errors.New("fetching tmdb err")
	}
	var response struct {
		Results []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			GenreIDs    []int  `json:"genre_ids"`
			ReleaseDate string `json:"release_date"`
			Overview    string `json:"overview"`
			Poster string 		`json:"poster_path"`
			} `json:"results"`
		}
		if err := json.Unmarshal(data, &response); err != nil {
			return turso.Content{}, err
		}

		if len(response.Results) == 0 {
			return turso.Content{}, errors.New("no data found for this movie")
		}
		show := response.Results[0]
		content := turso.Content{
			ID:          show.ID,
			Title:       show.Title,
			Genres:      mapGenreIDToString(show.GenreIDs),
			ReleaseDate: show.ReleaseDate,
		Description: util.TranslateToDutch(show.Overview),
		Provider: provider,
		ImageLink:   show.Poster,
		Type:        "movie",
	}
	
	return content, nil
}



func FetchGenresMovie() error {
	url := fmt.Sprintf("%s/genre/movie/list?language=en", BaseURL)
	data, err := makeRequest(url)
	if err != nil {
		return err
	}

	var response struct {
		Genres []Genre `json:"genres"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	for _, genre := range response.Genres {
		genreMap[genre.ID] = genre.Name
	}
	return nil
}

func FetchGenresShow() error {
	url := fmt.Sprintf("%s/genre/show/list?language=en", BaseURL)
	data, err := makeRequest(url)
	if err != nil {
		return err
	}

	var response struct {
		Genres []Genre `json:"genres"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	for _, genre := range response.Genres {
		genreMap[genre.ID] = genre.Name
	}
	return nil
}