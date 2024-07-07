package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	// Visit the link
	// 	e.Request.Visit(e.Attr("href"))

	// 	// Find img tags and get alt attribute
	// 	e.ForEach("img", func(_ int, img *colly.HTMLElement) {
	// 		alt := img.Attr("alt")
	// 		fmt.Printf("Image Alt: %s\n", alt)
	// 	})
	// })

	// Find img tags and get alt attribute
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// Find img tags and get alt attribute
		e.ForEach("img", func(_ int, img *colly.HTMLElement) {
			alt := img.Attr("alt")
			fmt.Printf("Image Alt: %s\n", alt)
		})
	})

	// Log visited URLs
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Start scraping from the given URL
	c.Visit("https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5")
}
