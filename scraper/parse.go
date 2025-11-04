package scraper

import (

	"github.com/PuerkitoBio/goquery"
)

func Scrape(url string) Result {
	resp, err := Fetch(url)

	if err != nil {
		return Result{
			URL:   url,
			Error: err,
		}
	}

	defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return Result{
				URL:   url,
				Error: err,
			}
		}

		title := doc.Find("title").Text()
		description := doc.Find("meta[name='description']").AttrOr("content", "")

		return Result{
			URL:         url,
			Title:       title,
			Description: description,
			Error:       nil,
		}

	
}
