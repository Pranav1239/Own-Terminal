package scraper

import (
	"fmt"

	"github.com/Pranav1239/Own-Terminal/pkg/utils"
	"github.com/gocolly/colly"
)

func MainScrapper(scrap string) error {
	switch scrap {
	case "1":
		return ProductScrapper()
	default:
		return fmt.Errorf("invalid scrapper option: %s", scrap)
	}
}

func ProductScrapper() error {
	fmt.Println("Hello, World!")

	// Example usage of Colly for web scraping
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title found:", e.Text)
	})

	err := c.Visit("http://example.com")
	if err != nil {
		utils.CheckErr(err)
		return err
	}

	return nil
}
