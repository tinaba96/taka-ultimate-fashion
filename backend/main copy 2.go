package main

import (
	"fmt"

	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/sclevine/agouti"
)

func main() {
	// c := colly.NewCollector()

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

	
	// JavaScriptを使った動的ページのHTMLを取得するためにChromeを利用することを宣言
    agoutiDriver := agouti.ChromeDriver(
        agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox", "--disable-dev-shm-usage"}),
    )
	// agoutiDriver := agouti.ChromeDriver()

    if err := agoutiDriver.Start(); err != nil {
        fmt.Println("Failed to start driver:%v", err)
        // return c.JSON(500, map[string]interface{}{"error": "agoutiDriver.Start()でエラー発生"})
    }

	defer agoutiDriver.Stop()
    // page, _ := agoutiDriver.NewPage()
    page, err := agoutiDriver.NewPage()
	if err != nil {
        fmt.Println("NewPage()でエラー発生")
        fmt.Println("Some context: %v", err)
        // return c.JSON(500, map[string]interface{}{"error": "NewPage()でエラー発生"})
    }



	// 各劇場サイトに入る
	err = page.Navigate("https://hlo.tohotheater.jp")
	// err = page.Navigate("https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5")
	if err != nil {
		fmt.Println("Failed to Navigate: ", err)
	}

	fmt.Println("プログラム開始")
	// 動的ページのHTMLを格納
	// dom, _ := page.HTML()
	dom, err := page.HTML()

	if err != nil {
		fmt.Println("page.HTML()でエラー発生")
		fmt.Println("Some context: %v", err)
	}

	fmt.Println("処理中")
	// スクレイピングするためにDOMを読み込みなおす
	contents := strings.NewReader(dom)


	doc, _ := goquery.NewDocumentFromReader(contents)
	// doc, err := goquery.NewDocumentFromReader(contents)
        // if err != nil {
        //     fmt.Println("page.HTML()でエラー発生")
        //     fmt.Println("Some context: %v", err)
        // }
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("alt")
		// fmt.Println(url)
		fmt.Printf("ALT : %s\n", url)
	})
	fmt.Println("処理完了")


	// // Find img tags and get alt attribute
	// c.OnHTML("html", func(e *colly.HTMLElement) {
	// 	// Find img tags and get alt attribute
	// 	e.ForEach("img", func(_ int, img *colly.HTMLElement) {
	// 		alt := img.Attr("alt")
	// 		fmt.Printf("Image Alt: %s\n", alt)
	// 	})
	// })

	// // Log visited URLs
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })




	// Start scraping from the given URL
	// c.Visit("https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5")
}
