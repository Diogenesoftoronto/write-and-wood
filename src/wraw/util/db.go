package util

import (
	"os"
	"log"
	"strings"
	"github.com/joho/godotenv"
	"github.com/charmbracelet/charm/kv"
	"github.com/dgraph-io/badger/v3"
)

// Open a database (or create one if it doesn’t exist)
db, err := kv.OpenWithDefaults("wraw-cli-db")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Fetch updates and easily define your own syncing strategy
if err := db.Sync(); err != nil {
    log.Fatal(err)
}