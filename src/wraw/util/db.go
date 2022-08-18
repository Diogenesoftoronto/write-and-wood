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




// Quickly get a value
// v, err := db.Get([]byte("dog"))
// if err != nil {
// 	panic(err)
// }

// Save some data
// if err := db.Set([]byte("fave-food"), []byte("gherkin")); err != nil {
//     log.Fatal(err)
// }

// All data is binary
// if err := db.Set([]byte("profile-pic"), someJPEG); err != nil {
//     log.Fatal(err)
// }

// Go full-blown Badger and use transactions to list values and keys
// db.View(func(txn *badger.Txn) error {
// 	opts := badger.DefaultIteratorOptions
// 	opts.PrefetchSize = 10
// 	it := txn.NewIterator(opts)
// 	defer it.Close() //nolint:errcheck
// 	for it.Rewind(); it.Valid(); it.Next() {
// 		item := it.Item()
// 		k := item.Key()
// 		err := item.Value(func(v []byte) error {
// 			fmt.Printf("%s - %s\n", k, v)
// 			return nil
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return nil
// })