package database;
import (
	"database/sql"
  "fmt"
  "os"
  "path/filepath"

  "github.com/tursodatabase/go-libsql"
)
func RunTursoDB(){
	dbName := "local.db"
    primaryUrl := "https://embeddings-realrubr2.turso.io"
    authToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJxVHFEVk9KR0VlLWVqeG9vV1FBbTFRIn0.OQ1MLWN8ztbZQPe5H31E1tc3PAAFH6hBCTqqo7g_39SQfURH47qA2rHObsg64j75KcVhCMU83U6Ko7Jg7DTlAw"

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