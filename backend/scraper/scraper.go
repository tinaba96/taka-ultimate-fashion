package scraper

// #cgo CFLAGS: -I/Library/Frameworks/Python.framework/Versions/3.9/include/python3.9
// #cgo LDFLAGS: -L/Library/Frameworks/Python.framework/Versions/3.9/lib -lpython3.9
// #include <Python.h>

import (
	"C"
	"os"
	"strconv"

	"gorm.io/gorm"
)
import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"gorm.io/driver/mysql"
)

type Product struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string
    Price      string
    ImageURL   string
	CategoryID uint
	Category   Category
    // Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
    gorm.Model
}

type Category struct {
    ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
    gorm.Model
}

type Save struct {
    ID   uint   `gorm:"primaryKey"`
	ProductID uint
	Product Product
	Status uint
    gorm.Model
}

type DBConfig struct {
	User string
	Password string
	Host string
	Port int
	Table string
}

func getDBConfig() DBConfig {
    port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
    return DBConfig{
        User: os.Getenv("MYSQL_USER"),
        Password: os.Getenv("MYSQL_PASSWORD"),
        Host: os.Getenv("MYSQL_HOST"),
        Port: port,
		Table: os.Getenv("MYSQL_DB"),
    }
}

func Scraper() {

	// ====== DB Connections =======
	// https://qiita.com/Alphard/items/5239c2707b1557f76c9a#maingo
	config := getDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	// dsn := "user:password@tcp(db:3306)/taka_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errdb != nil {
		panic("failed to connect database")
	}

	// var product Product
	// db.First(&product, "Name = ?", "TSHIRT") // find product with code D42
	// fmt.Println(product)
	
	log.Println("start scraping.....")

	// ====== start scraping =======
	url := "https://oldnavy.gapcanada.ca/browse/category.do?cid=5249&mlink=5155,1,m_5" // specify the URL
	// url := "https://oldnavy.gapcanada.ca/browse/category.do?cid=5199&mlink=5155,1,m_6"

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
	

	contentsDom.Find(".cat_product-image").Each(func(i int, s *goquery.Selection) {
		imgSrc, _ := s.Find("img").Attr("src")
		productName := s.Next().Find("div.category-page-1r1wcud").Text()
		price := s.Next().Next().Find("div.product-price__highlight").Text()

		var productCategory Category
		db.FirstOrCreate(&productCategory, Category{Name: "T-SHIRTS"}) 
		// db.FirstOrCreate(&productCategory, Category{Name: "JEANS"}) 

		product := Product{
			Name:     productName,
			Price:    price,
			ImageURL: imgSrc,
			CategoryID: productCategory.ID,
		}

		if err := db.Create(&product).Error; err != nil {
			log.Printf("Failed to save product: %v", err)
		} else {
			fmt.Printf("Saved Product: %s, Price: %s, Image Src: %s\n", productName, price, imgSrc)
		}

		fmt.Printf("Product: %s, Price: %s, Image Src: %s\n", productName, price, imgSrc)
	})

	log.Println("Scraping Completed")
	
}