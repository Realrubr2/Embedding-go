package tmdb
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"log"
	"os"
	"strings"
	"embeddings/turso"
)

const APIKey = ""

func convertMovieToContent(movie Movie) turso.Content {
	// Convert genre slice to a single comma-separated string
	genreString := strings.Join(movie.Genres, ", ")

	// Convert `Movie` to `Content`
	content := turso.Content{
		ID:          movie.ID,
		Title:       movie.Title,
		Genres:      genreString, // Store as a single string
		Description: movie.Description,
		ImageLink:   movie.ImageLink,
		ReleaseDate: movie.ReleaseDate,
		Provider:    movie.Provider,
		Type:        "movie", // Explicitly setting the type
	}
	return content
}
//maps genreid to string it updates the mapgenreidstonames, it really calls for a refactor...
func mapGenreIDToString(ids []int)string{
	var names []string
	for _, id := range ids {
		if name, exists := genreMap[id]; exists {
			names = append(names, name)
		}
	}
	genreString := strings.Join(names, ", ")
	return genreString
}

// Maps genre IDs to genre names
func mapGenreIDsToNames(ids []int) []string {
	var names []string
	for _, id := range ids {
		if name, exists := genreMap[id]; exists {
			names = append(names, name)
		}
	}
	return names
}

// Makes an HTTP GET request and returns the response body
func makeRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+APIKey)
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// Check for HTTP errors (non-2xx status codes)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status: %d %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	return io.ReadAll(res.Body)
}


func WriteToFile(fileName string, object []Movie){
fileNameString := fmt.Sprintf("%s.json",fileName)
	// Open a file for writing
	file, err := os.Create(fileNameString)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	

	// Encode the structured JSON with indentation
	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	if err := enc.Encode(object); err != nil {
		log.Fatal("Error encoding JSON: ", err)
	}

	fmt.Println("JSON successfully saved to json")

}