package main

import (
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

	fmt.Printf("RSS feed title: %s\n", feed.Title)
	fmt.Printf("RSS feed description: %s\n", feed.Description)
	fmt.Printf("RSS feed link: %s\n", feed.Link)
	fmt.Printf("RSS feedlink: %s\n", feed.FeedLink)
	fmt.Printf("RSS last updated on: %s\n", feed.Updated)
	fmt.Printf("-----------------------------------------------------\n")
	for index, feedItem := range feed.Items {
		fmt.Printf("Item %d\n", index+1)
		fmt.Printf("Link: %s\n", feedItem.Title)
		fmt.Printf("Published date: %s\n", feedItem.Published)
		fmt.Printf("-----------------------------------------------------\n")
	}
}
