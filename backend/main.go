package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/sclevine/agouti"
)

func main() {
	url := "https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5" // specify the URL

	// driver := agouti.ChromeDriver()                                // Start the driver
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-gpu",
			"--whitelisted-ips",
			"--detach",
		}),
		// agouti.Debug,  // Local only
	)

	log.Println("start")
	err := driver.Start()
	if err != nil {
		log.Printf("Error starting driver: %v", err.Error())
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome")) // Start Chrome and return a page type value (session)
	if err != nil {
		log.Printf("Error creating new page: %v", err.Error())
	}
	time.Sleep(3 * time.Second)

	err = page.Navigate(url) // Access the specified URL
	if err != nil {
		log.Printf("Error navigating to job post link: %v", err.Error())
	}
	time.Sleep(3 * time.Second)
	log.Println("end")

	// Scroll down to load more content
	for i := 0; i < 5; i++ {
		err = page.RunScript("window.scrollTo(0, document.body.scrollHeight);", nil, nil)
		if err != nil {
			log.Fatalf("Failed to scroll: %v", err)
		}
		time.Sleep(3 * time.Second) // Wait for new content to load
	}

	log.Println("Retrieving HTML content...")

	curContentsDom, err := page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	// Write the HTML to output.txt
	err = os.WriteFile("output.txt", []byte(curContentsDom), 0644)
	if err != nil {
		log.Printf("Failed to write to file: %v", err)
		return
	}
	log.Println("HTML content written to output.txt")

	readerCurContents := strings.NewReader(curContentsDom)
	contentsDom, _ := goquery.NewDocumentFromReader(readerCurContents) // Get the DOM of the currently open page in the browser
	// listDom := contentsDom.Find("ul").Children()                       // selector - Enter the selector for the select box in the selector part. Get all the child elements in the select box
	// log.Println("[SELECTOR]", listDom)
	
	contentsDom.Find("div img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("alt")
		// fmt.Println(url)
		fmt.Printf("ALT : %s\n", url)
	})
	contentsDom.Find("div.category-page-1r1wcud").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		fmt.Printf("Text: %s\n", text)
	})
}
