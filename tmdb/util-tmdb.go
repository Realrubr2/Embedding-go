package tmdb

import (
	"embeddings/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// const APIKey = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJjMzkyM2U1YmY3MzhhMTc5YTNlMzk3MjUzYjM0NmJkZSIsIm5iZiI6MTczNzIzMDQxOC43NTYsInN1YiI6IjY3OGMwODUyOGJjYTY2MWQwNTQzMWE2OSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ioPqDhNGjeyIBThSJErhQ1JXJO1cnKn4jTfmoSnTXyA"

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


// Makes an HTTP GET request and returns the response body
func makeRequest(url string) ([]byte, error) {
	env :=util.LoadEnviroment()
	fmt.Sprintln(env[2])
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+env[1])
	
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