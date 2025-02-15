package tmdb

import (
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