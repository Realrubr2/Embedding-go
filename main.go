package main

import (
	"database/sql"
	"embeddings/scrape"
	"embeddings/util"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

    env := util.LoadEnviroment()
    primaryUrl := env[3]
    authToken := env[2]
    url := fmt.Sprintf("%s?authToken=%s", primaryUrl, authToken)
    db, err := sql.Open("libsql", url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
      os.Exit(1)
    }

    scrape.ScrapeAll(db)
   
    defer db.Close()
}

