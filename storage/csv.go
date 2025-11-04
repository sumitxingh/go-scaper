package storage

import (
	"encoding/csv"
	"go-scraper/scraper"
	"os"
)

func SaveToCSV(filename string, results []scraper.Result) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, d := range results {
		writer.Write([]string{d.URL, d.Title, d.Description})
	}

	return nil

}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
