package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
	"github.com/mohae/struct2csv"
)

func main() {

	rssFeeds := []string{"https://akashgoswami.com/articles/index.xml", "https://bradleyjkemp.dev/index.xml"}
	feeds := []gofeed.Feed{}

	for _, v := range rssFeeds {
		feedParser := gofeed.NewParser()
		feed, err := feedParser.ParseURL(v)
		if err != nil {
			log.Fatalf("Error getting feed: %v", err)
		}
		feeds = append(feeds, *feed)
	}

	//fmt.Printf("Feeds: %v", feeds)

	enc := struct2csv.New()
	rows, err := enc.Marshal(feeds)
	if err != nil {
		log.Fatalf("Error encoding data: %v", err)
	}

	csvFile, err := os.Create("rss.csv")
	if err != nil {
		log.Fatalf("Error creating csv file: %v", err)
	}
	defer csvFile.Close()

	write := csv.NewWriter(csvFile)
	err = write.WriteAll(rows)
	if err != nil {
		log.Fatalf("Error writing data to csv file: %v", err)
	}

	// Mashaling a feed into JSON
	// jsonData, err := json.Marshal(feed)
	// if err != nil {
	// 	log.Fatalf("Error marshaling data to JSON: %v", err)
	// }
	// fmt.Println("Struct as JSON: ", string(jsonData))
}
