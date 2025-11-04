package scraper

import (
	"net/http"
	"time"
)

func Fetch(url string) (*http.Response, error) {

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "GoScraperBot/1.0")
	return client.Do(req)

}