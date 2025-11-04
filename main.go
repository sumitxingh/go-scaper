package main

import (
	"fmt"
	"go-scraper/scraper"
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
	}

	results := scraper.Run(urls, 10)

	for _, result := range results {
		fmt.Println(result.URL, result.Title, result.Description)
	}
}