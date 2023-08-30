package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func main() {

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL("https://akashgoswami.com/articles/index.xml")
	if err != nil {
		log.Fatalf("Error getting feed: %v", err)
	}

	jsonData, err := json.Marshal(feed)
	if err != nil {
		log.Fatalf("Error marshaling data to JSON: %v", err)
	}
	fmt.Println("Struct as JSON: ", string(jsonData))
}
