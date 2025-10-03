package main

import (
    "fmt"
    "encoding/xml"
    "net/http"
    "io/ioutil"
)

type RSS struct {
    Channel Channel `xml:"channel"`
}
type Channel struct {
    Title string `xml:"title"`
    Items []Item `xml:"item"`
}
type Item struct {
    Title string `xml:"title"`
    Link  string `xml:"link"`
}
func fetchRSS(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	err = xml.Unmarshal(bytes, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}
func main() {
    feedUrl:= "https://blog.golang.org/feed.atom"
    rss,err:= fetchRSS(feedUrl)
    if err != nil {
        fmt.Println("Error fetching RSS feed:", err)
        return
    }
    	fmt.Println("Feed Title:", rss.Channel.Title)
	for _, item := range rss.Channel.Items {
		fmt.Printf("- %s (%s)\n", item.Title, item.Link)
	}
}
