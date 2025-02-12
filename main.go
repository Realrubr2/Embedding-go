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
    authToken := ""

    url := fmt.Sprintf("%s?authToken=%s", primaryUrl, authToken)
    db, err := sql.Open("libsql", url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
      os.Exit(1)
    }

	  // tmdb.Movies(db)
    tmdb.Shows(db)
    // util.TranslateDescription()
  
    defer db.Close()
}
