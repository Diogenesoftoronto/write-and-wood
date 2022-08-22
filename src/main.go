package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/charm/kv"
)

// loop through values in the json and set them in the db

type Answer struct {
	id      string `json:"Id"`
	content string `json:"Content"`
	created string `json:"Created"`
}

type Engagement struct {
	Rating string `json:"Rating"`
	Views  string `json:"Views"`
	Age    string `json:"Age"`
	Total  string `json:"Total"`
}

type PromptData struct {
	Id          int        `json:"Id"`
	Content     string     `json:"Content"`
	Created     string     `json:"Created"`
	Type        string     `json:"Type"`
	Author      string     `json:"Author"`
	Answers     []Answer   `json:"Answers"`
	Engagements Engagement `json:"Engagement"`
}

func main() {
	// Open a database (or create one if it doesn’t exist)
	db, err := kv.OpenWithDefaults("wraw-cli-db")
	// fmt.Println("type of db: ", reflect.TypeOf(db))
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

	for k, v := range prompts {
		fmt.Printf("adding PromptData %v to database. Value: %v\n", k, v)
		// encode promptdata into a byte array
		jsonString := fmt.Sprintf(`{
			"Content": "%v",
			"Id": %v,
			"Type": "%v",
			"Author": "NONE",
			"Created": "%s",
			"Answers": [],
			"Engagment": {
				"Rating": 0,
				"Views": 0,
				"Age": 0,
				"Total": 0
			}
		}`, v.Content, strconv.Itoa(k), v.Type, time.Now())

		// set data in the data base
		if err := db.Set([]byte("PromptData "+strconv.Itoa(k)), []byte(jsonString)); err != nil {
			log.Fatal("data was not set in db: ", err)
		}
		fmt.Printf("Succesfully added PromptData %v to database.\n", k)

	}
	// get values from the db
	for k, v := range prompts {
		// get data from the data base
		fmt.Printf("getting data from database: %v\n", v)
		data, err := db.Get([]byte("PromptData " + strconv.Itoa(k)))
		if err != nil {
			log.Fatal("error getting data from db: ", err)
		}
		// convert the data into a promptdata struct
		var data_decoded PromptData
		if err = json.Unmarshal(data, &data_decoded); err != nil {
			log.Fatal("error unmarshalling data: ", err, "| data: ", string(data))
		}

		fmt.Printf("data retreived successfully from database: %s\n", data_decoded)
	}

}
