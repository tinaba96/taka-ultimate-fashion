package model

import (
	"github.com/tinaba96/taka-ultimate-fashion/backend/database"
	"gorm.io/gorm"
)

// type Product struct {
// 	ID       uint `gorm:"primaryKey"`
// 	Name     string
// 	Price    string
// 	ImageURL string
// 	gorm.Model
// }
// type Category struct {
//     ID   uint   `gorm:"primaryKey"`
//     Name string
//     gorm.Model
// }

type Save struct {
    ID   uint   `gorm:"primaryKey"`
	ProductID uint
	Product Product
	Status uint
    gorm.Model
}

type RequestBodyType struct {
	ProductID uint   `json:"productID" binding:"required"`
	// Status    uint   `json:"status" binding:"required"`
	Status   uint 
}

func SaveProducsst() (products []Product) {
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
	return
	// return result
}

func GetAllSaveProduct() (saves []Save) {
	database.Init(false)
	Db := database.GetDB()
	result := Db.Preload("Product").Find(&saves)
	// result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}


func PostProduct(requestBody RequestBodyType) (error) {
	db := database.GetDB()

	saved := Save{
		ProductID: requestBody.ProductID,
		Status:    requestBody.Status,
	}

	// if err := db.Create(&saved).Error; err != nil {
	// 	return err
	// }

	// var saved Save
	result := db.Where("product_id = ?", requestBody.ProductID).First(&saved)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.RowsAffected == 0 {
		// saved = Save{
		// 	ProductID: requestBody.ProductID,
		// 	Status:    requestBody.Status,
		// }
		if err := db.Create(&saved).Error; err != nil {
			return err
		}
	} else {
		saved.Status = requestBody.Status
		if err := db.Save(&saved).Error; err != nil {
			return err
		}
	}

	return nil

}