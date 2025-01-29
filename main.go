package main

import (
	"database/sql"
	 "embeddings/tmdb"
	// "embeddings/util"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)



func main() {

	// dbName := "local.db"
    primaryUrl := "libsql://embeddings-realrubr2.turso.io"
    authToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJleHAiOjE3Mzg0MTI2NjYsImlhdCI6MTczNzgwNzg2NiwiaWQiOiI2MGEzZDY3Ni1mZTA5LTRhYmYtYjY2OS0wNDhmMjI4ZTZjYTMifQ.OgooF_Kz3x2w6ZkV6Ge3ZPhEYcbSihZI5JfC6QMQl4DktotdpZGFwRnmmKz1HiCq_YCgV8VwxI0w8t42ue_uAQ"

    url := fmt.Sprintf("%s?authToken=%s", primaryUrl, authToken)
    db, err := sql.Open("libsql", url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
      os.Exit(1)
    }

	  repo.Movies(db)
    repo.Shows(db)
    // util.TranslateDescription()
  
    defer db.Close()
}
