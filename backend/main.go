package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/PuerkitoBio/goquery"

	"github.com/sclevine/agouti"
)
type Product struct {
	gorm.Model
	Name  string
	Price uint
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


func main() {
	// ====== DB Connections =======
	config := getDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	// dsn := "user:password@tcp(db:3306)/taka_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errdb != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Name: "TSHIRT", Price: 100})

	var product Product
	db.First(&product, "Name = ?", "TSHIRT") // find product with code D42
	fmt.Println(product)



	// ====== start scraping =======
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
