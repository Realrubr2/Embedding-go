package turso

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)
type Content struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Genres      string 	`json:"genres"`
	ReleaseDate string   `json:"release_date"`
	Description string   `json:"overview"`
	Provider    string   `json:"provider,omitempty"`
	ImageLink string	 `json:"image_Path"`
	Type string			`json:"type"`
}

type Embeddings struct {
	Content_ID int
	Vector []float64
}

func TurnIntoVector(){}

func GetContent(db *sql.DB){
		rows, err := db.Query("SELECT * FROM content")
		if err != nil {
		  fmt.Fprintf(os.Stderr, "failed to execute query: %v", err)
		  os.Exit(1)
		}
		defer rows.Close()
	  
		// var contents []content
	  
		for rows.Next() {
		  var content Content
	  
		  if err := rows.Scan(&content.ID,&content.Title,&content.Genres,&content.Description,&content.ImageLink,&content.ReleaseDate,&content.Provider,&content.Type); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		  }
	  
		//   contents = append(contents, content)
		  fmt.Println(content.ID, content.Title)
		}
	  
		if err := rows.Err(); err != nil {
		  fmt.Println("Error during rows iteration:", err)
		}
}

func CreateContent(db *sql.DB, content Content) {
	_, err := db.Exec(`
		INSERT OR REPLACE INTO content (id, title, genres, description, image_link, release_date, provider, type) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		content.ID, content.Title, content.Genres, content.Description, content.ImageLink, content.ReleaseDate, content.Provider, content.Type)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}


	fmt.Println("Content inserted successfully:", content.ID)
}

// func CreateEmbeddings(db *sql.DB, embeddingObj Embeddings){

// 	embeddingJSON, err := json.Marshal(embeddingObj.Vector)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "failed to convert vector to JSON: %v\n", err)
// 		os.Exit(1)
// 	}

// 	result, err := db.Exec("INSERT INTO embeddings(content_id, vectors) VALUES(?, ?)", embeddingObj.Content_ID, embeddingJSON)
// 	if err != nil {
// 	  fmt.Fprintf(os.Stderr, "failed to execute query: %v", err)
// 	  os.Exit(1)
// 	}
// 	rows, _ := result.RowsAffected()
// 	if rows > 1 {
// 		fmt.Println(" sucsess", rows)
// 	}
// }
func CreateEmbeddings(db *sql.DB, embeddingObj Embeddings) {
	// Ensure that the vector length is 1536
	if len(embeddingObj.Vector) != 1536 {
		fmt.Fprintf(os.Stderr, "Error: embedding vector must have exactly 1536 dimensions, but got %d\n", len(embeddingObj.Vector))
		os.Exit(1)
	}
	vectorStr, err := json.Marshal(embeddingObj.Vector);
	if err != nil {
		log.Fatal("error encoding json")
	}
	// Convert the vector to a string format compatible with vector32 (for float64)

	// Insert the data into the database using vector32
	result, err := db.Exec("INSERT OR REPLACE INTO embeddings(content_id, vectors) VALUES(?, vector32(?))", embeddingObj.Content_ID, string(vectorStr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}

	// Check if the insert was successful
	rows, _ := result.RowsAffected()
	if rows > 0 {
		fmt.Println("Success: Embedding inserted successfully")
	} else {
		fmt.Println("Failed to insert embedding")
	}
}
