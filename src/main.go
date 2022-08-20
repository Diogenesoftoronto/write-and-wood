package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/charmbracelet/charm/kv"
)

// loop through values in the json and set them in the db

// type Answer struct {
// 	id      string `json:"Id"`
// 	content string `json:"Content"`
// 	created string `json:"Created"`
// }

type Engagement struct {
	Rating string `json:"Rating"`
	Views  string `json:"Views"`
	Age    string `json:"Age"`
}

type PromptData struct {
	Id          int        `json:"Id"`
	Content     string     `json:"Content"`
	Created     string     `json:"Created"`
	Type        string     `json:"Type"`
	Author      string     `json:"Author"`
	Answers     []string   `json:"Answers"`
	Engagements Engagement `json:"Engagement"`
}

func main() {
	// Open a database (or create one if it doesn’t exist)
	db, err := kv.OpenWithDefaults("wraw-cli-db")
	fmt.Println("type of db: ", reflect.TypeOf(db))
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Fetch updates and easily define your own syncing strategy
	if err := db.Sync(); err != nil {
		log.Fatal("error syncing database: ", err)
	}

	// unmarshal the json file into the db
	jsonFIle, err := os.Open("all-prompts.json")
	if err != nil {
		log.Fatal("error opening json file: ", err)
	}
	defer jsonFIle.Close()
	byteJson, _ := ioutil.ReadAll(jsonFIle)
	var prompts []PromptData

	if err := json.Unmarshal(byteJson, &prompts); err != nil {
		log.Fatal("error unmarshalling json: ", err)
	}
	// func(db *kv.KV, values []PromptData) {
	//     for k, v := range prompts {
	// 	// err := db.Update(func(txn *badger.Txn) error {
	// 	//     return txn.Set([]byte(k), []byte(v))
	// 	// })
	// 	if err := db.Set(k, v); err != nil {
	// 		log.Fatal("")
	// 	}

	//     }
	// }

	for k, v := range prompts {
		// encode promptdata into a byte array
		var promptbuf bytes.Buffer
		encoder := gob.NewEncoder(&promptbuf)
		fmt.Println("prompts added to db %s, %s", k, v)
		if err := encoder.Encode(v); err != nil {
			log.Fatal("error encoding promptdata: ", err)
		}
		// set data in the data base
		if err := db.Set([]byte(strconv.Itoa(k)), promptbuf.Bytes()); err != nil {
			log.Fatal("")
		}

	}

	// setValues(&db, prompts)
	// db.Get([]byte("1"))
	// for k, v := range prompts {

	// 	fmt.Println("prompts added to db %s, %s", k, reflect.TypeOf(v).Kind())
	// }

}
