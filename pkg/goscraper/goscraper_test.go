package goscraper

import (
	"fmt"
	"testing"
)

func TestScrape(t *testing.T) {
	scrape, err := Scrape("https://www.baidu.com", 1)
	if err != nil {
		return
	}
	fmt.Println(scrape)
}
