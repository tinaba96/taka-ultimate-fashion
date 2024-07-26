package main

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

	"github.com/tinaba96/taka-ultimate-fashion/backend/router"
	"github.com/tinaba96/taka-ultimate-fashion/backend/scraper"
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

func main() {
	// ====== API =======
	log.Println("start server...")
	router := router.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}


	// ====== Python Execution =======
	// cmd := exec.Command("python3", "deep_learning/keras.py")

    // // コマンドを実行し、標準出力を取得
    // output, err := cmd.CombinedOutput()
    // if err != nil {
    //     fmt.Println("Error:", err)
    //     return
    // }
	// fmt.Println(string(output))
	// err := python.PyRun_SimpleFile("./deep_learning/keras.py")
    // if err != nil {
    //     panic(err.Error)
    // }

	

	// ====== DB Connections =======
	// https://qiita.com/Alphard/items/5239c2707b1557f76c9a#maingo
	config := getDBConfig();
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Table)
	// dsn := "user:password@tcp(db:3306)/taka_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, errdb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errdb != nil {
		panic("failed to connect database")
	}


	// Migrate the schema
	// if !db.Migrator().HasTable(&Product{}) && !db.Migrator().HasTable(&Category{}) && !db.Migrator().HasTable(&Save{}) {
	// db.AutoMigrate(&Product{}, &Category{}, &Save{})
		// Master Category Data
	db.FirstOrCreate(&Category{Name: "T-SHIRTS"})
	db.FirstOrCreate(&Category{Name: "SWIMWEAR"})
	db.FirstOrCreate(&Category{Name: "SHIRTS"})
	db.FirstOrCreate(&Category{Name: "JEANS"})
	db.FirstOrCreate(&Category{Name: "SHORTS"})
	db.FirstOrCreate(&Category{Name: "PANTS"})
	db.FirstOrCreate(&Category{Name: "GRAPHIC T-SHIRTS"})
	db.FirstOrCreate(&Category{Name: "ACTIVEWEAR"})
	// }


	scraper.Scraper()
		
	
	// db.FirstOrCreate(&Save{ProductID: 12, Status: 1}) 
	log.Println("--- DONE ---")

}