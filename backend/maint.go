package main

import (
	// "github.com/gin-gonic/gin"
	"fmt"

	// "io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main2() {
	fmt.Println("Hello World")
	// 1. ウェブページの取得
	// res, err := http.Get("https://oldnavy.gapcanada.ca/browse/product.do?pid=419283003&cid=1010005&pcid=1010005&vid=1#pdp-page-content")
	res, err := http.Get("https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5")
	// res, err := http.Get("https://zenn.dev/books/explore")
	// res, err := http.Get("https://oldnavy.gapcanada.ca/browse/division.do?cid=5155&mlink=5151%2CtopNav%2Cvisnav&nav=meganav%3AMen%3A%3A")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// 2. HTTPレスポンスのボディを読み取る
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }



	// // HTML全体を文字列として表示
	// htmlString := string(body)
	// println(htmlString)

	// // テキストファイルに出力する
	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// _, err = io.WriteString(file, htmlString)
	// if err != nil {
	// 	panic(err)
	// }




	// // 3. 取得したHTML全体を文字列として表示
	// fmt.Println(string(body))

	// 2. HTMLの解析
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// // 3. 必要な情報の抽出
	// title := doc.Find("title").Text()
	// fmt.Println(title)


	// 3. img要素の抽出
	// doc.Find("div div main dev dev dev section div div div div a img").Each(func(index int, item *goquery.Selection) {
	// 	// imgSrc, exists := item.Attr("src")
	// 	// if exists {
	// 	// 	fmt.Printf("Image %d: %s\n", index+1, imgSrc)
	// 	// }
	// 	imgAlt, exists := item.Attr("alt")
	// 	if exists {
	// 		fmt.Printf("Alt Text %d: %s\n", index+1, imgAlt)
	// 	}
	// })
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("alt")
		// fmt.Println(url)
		fmt.Printf("URL : %s\n", url)
   })


	doc.Find(".category-page-1r1wcud").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println("Found:", text)
	})


}
