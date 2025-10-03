package main

import (
	"fmt"
)

func main() {
	feedURL := "https://hnrss.org/frontpage"
	rss, err := fetchRSS(feedURL)
	if err != nil {
		fmt.Println("Error fetching RSS feed:", err)
		return
	}
	fmt.Println("Feed Title:", rss.Channel.Title)
	for _, item := range rss.Channel.Items {
		fmt.Printf("- %s (%s)\n", item.Title, item.Link)
	}
}
