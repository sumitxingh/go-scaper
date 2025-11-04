package main

import (
	"fmt"
	"go-scraper/scraper"
	"go-scraper/storage"
)


func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.pinterest.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.ask.com",
		"https://www.ycombinator.com",
		"https://www.crunchbase.com",
		"https://www.angellist.com",
		"https://www.producthunt.com",
		"https://www.hackernews.com",
		"https://www.news.ycombinator.com",
	}

	results := scraper.Run(urls, 10)

	err := storage.SaveToCSV("results.csv", results)
	if err != nil {
		fmt.Println("Error saving to CSV:", err)
	}

	for _, result := range results {
		fmt.Println(result.URL, result.Title, result.Description)
	}
}