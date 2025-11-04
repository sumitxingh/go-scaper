package scraper

import "sync"

type Result struct {
	URL         string
	Title       string
	Description string
	Error       error
}

func Run(urls []string, concurrency int) []Result {
	jobs := make(chan string, len(urls))
	result := make(chan Result, len(urls))

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				res := Scrape(url)
				result <- res
			}
		}()
	}

	for _, u := range urls {
		jobs <- u
	}

	close(jobs)
	wg.Wait()
	close(result)

	var all []Result

	for r := range result {
		all = append(all, r)
	}

	return all
}
