package main

import (
	"fmt"
)
func main() {
	feeds:=[]string{
		"https://hnrss.org/newest",      
		"https://hnrss.org/frontpage",     
		"https://xkcd.com/rss.xml",  	
	}
	ch:=make(chan *RSS)
		for _, url := range feeds {
		go func(u string) {
			rss, err := fetchRSS(u)
			if err != nil {
				fmt.Println("Error fetching:", u, err)
				ch <- nil
				return
			}
			ch <- rss
		}(url)
	}
	for range feeds {
		rss := <-ch
		if rss != nil {
			fmt.Println("Feed:", rss.Channel.Title)
			for i, item := range rss.Channel.Items {
				if i >= 3 { 
					break
				}
				fmt.Printf("  - %s (%s)\n", item.Title, item.Link)
			}
			fmt.Println()
		}
	}
}
