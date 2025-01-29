package repo

import (
	"encoding/json"
	"fmt"

)

// Fetches popular movies from TMDB
func fetchPopularMovies() ([]Movie, error) {
	url := fmt.Sprintf("%s/movie/popular?language=en-US&page=1", BaseURL)
	data, err := makeRequest(url)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	var movies []Movie
	for _, r := range response.Results {
		movies = append(movies, Movie{
			ID:          r.ID,
			Title:       r.Title,
			Genres:      mapGenreIDsToNames(r.GenreIDs),
			ReleaseDate: r.ReleaseDate,
			Description: r.Overview,
			ImageLink: r.Poster,
		})
	}
	return movies, nil
}

func fetchPopularShows() ([]Content, error) {
	url := fmt.Sprintf("%s/tv/popular?language=en-US&page=1", BaseURL)
	data, err :=makeRequest(url)
	if err != nil {
		return nil, err
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
		return nil, err
	}


	var tvShows []Content
	for _, r := range response.Results {
		if r.Overview == "" { // Skip if Overview is empty
			r.Overview = "Sorry we dont have a desription for this show"
		}
		tvShows = append(tvShows, Content{
			ID:          r.ID,
			Title:       r.Title,
			Genres:      mapGenreIDToString(r.GenreIDs),
			ReleaseDate: r.ReleaseDate,
			Description: r.Overview,
			ImageLink:   r.Poster,
			Type:        "show",
		})
	}
	
return tvShows, nil
}


// Fetches streaming provider info for a given movie ID
func fetchMovieProvider(movieID int) (string, error) {
	url := fmt.Sprintf("%s/movie/%d/watch/providers", BaseURL, movieID)
	data, err := makeRequest(url)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		return "", err
	}

	if results, ok := response["results"].(map[string]interface{}); ok {
		if nl, found := results["NL"].(map[string]interface{}); found {
			if providers, exists := nl["flatrate"].([]interface{}); exists && len(providers) > 0 {
				if provider, valid := providers[0].(map[string]interface{}); valid {
					if name, exists := provider["provider_name"].(string); exists {
						return name, nil
					}
				}
			}
		}
	}

	return "Unknown", nil
}

func fetchShowProvider(showID int) (string, error) {
	url := fmt.Sprintf("%s/tv/%d/watch/providers", BaseURL, showID)
	data, err := makeRequest(url)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		return "", err
	}

	if results, ok := response["results"].(map[string]interface{}); ok {
		if nl, found := results["NL"].(map[string]interface{}); found {
			if providers, exists := nl["flatrate"].([]interface{}); exists && len(providers) > 0 {
				if provider, valid := providers[0].(map[string]interface{}); valid {
					if name, exists := provider["provider_name"].(string); exists {
						return name, nil
					}
				}
			}
		}
	}

	return "Unknown", nil
}

// Fetches and stores genre mappings from TMDB
func fetchGenresMovie() error {
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
func fetchGenresShow() error {
	url := fmt.Sprintf("%s/genre/tv/list?language=en", BaseURL)
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