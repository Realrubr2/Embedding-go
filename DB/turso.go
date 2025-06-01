package database

import (
	"database/sql"
	"embeddings/util"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func RunTursoDB(){
	dbName := "local.db"
    env := util.LoadEnviroment()
    primaryUrl := env[3]
    authToken := env[2]
    dir, err := os.MkdirTemp("", "libsql-*")
    if err != nil {
        fmt.Println("Error creating temporary directory:", err)
        os.Exit(1)
    }
    defer os.RemoveAll(dir)

    dbPath := filepath.Join(dir, dbName)

    connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
        libsql.WithAuthToken(authToken),
    )
    if err != nil {
        fmt.Println("Error creating connector:", err)
        os.Exit(1)
    }
    defer connector.Close()

    db := sql.OpenDB(connector)
    defer db.Close()
}