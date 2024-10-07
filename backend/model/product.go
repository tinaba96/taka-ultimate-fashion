package model

import (
	"fmt"
	"strconv"

	"github.com/tinaba96/taka-ultimate-fashion/backend/database"
	"gorm.io/gorm"
)

type Product struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Price    string `gorm:"unique"`
	ImageURL string
	gorm.Model
}
type Category struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
    gorm.Model
}


func GetAll() (products []Product) {
// func GetAll() []Product {
// 	var products []Product
// func GetAll() []gin.H {
	database.Init(false)
	// dsn := "user:password@tcp(db:3306)/taka_database?charset=utf8mb4&parseTime=True&loc=Local"
	// Db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db := database.GetDB()
// 	var result = []gin.H{
// 		{"id":      7791212,
// 		"name":     "iNew Basic Tee",
// 		"href":     "#",
// 		"imageSrc": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
// 		"imageAlt": "Front of men's Basic Tee in black.",
// 		"price":    "$35",
// 		"color":    "White",},
// 		{"id":       2,
// 		"name":     "Basic Tee",
// 		"href":     "#",
// 		"imageSrc": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
// 		"imageAlt": "Back of men's Basic Tee in black.",
// 		"price":    "$135",
// 		"color":    "Black",},
// }
	result := Db.Find(&products)
	// result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	database.Close()
	return
	// return result
}

func GetByCategories(categoryNames []string, minPrice string, maxPrice string) ([]Product, error) {
	db := database.GetDB()

	var categories []Category
	if err := db.Where("name IN ?", categoryNames).Find(&categories).Error; err != nil {
		return nil, err
	}

	if (len(categories) == 0 && minPrice == "" && maxPrice == "") {
		return nil, nil
	}

	var categoryIDs []uint
	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.ID)
	}

	minPriceInt, err := strconv.ParseFloat(minPrice, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid minPrice value: %w", err)
    }
	maxPriceInt, err := strconv.ParseFloat(maxPrice, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid maxPrice value: %w", err)
    }


	var products []Product
	if len(categoryIDs) == 0 {
	if err := db.Where("CAST(price AS DECIMAL(10,2)) BETWEEN ? AND ?", minPriceInt, maxPriceInt).Find(&products).Error; err != nil {
		return nil, err
	}
	}else{
	if err := db.Where("category_id IN ?", categoryIDs).Where("CAST(price AS DECIMAL(10,2)) BETWEEN ? AND ?", minPriceInt, maxPriceInt).Find(&products).Error; err != nil {
		return nil, err
	}
}
	database.Close()
	return products, nil
}

func GetByCostRange(categoryNames []string) ([]Product, error) {
	db := database.GetDB()

	var categories []Category
	if err := db.Where("name IN ?", categoryNames).Find(&categories).Error; err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, nil
	}

	var categoryIDs []uint
	for _, category := range categories {
		categoryIDs = append(categoryIDs, category.ID)
	}

	var products []Product
	if err := db.Where("category_id IN ?", categoryIDs).Find(&products).Error; err != nil {
		return nil, err
	}
	database.Close()
	return products, nil
}
