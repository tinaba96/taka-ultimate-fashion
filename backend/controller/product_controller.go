package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinaba96/taka-ultimate-fashion/backend/model"
)


func ShowAllProducts(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	categoryNames := c.DefaultQuery("categories", "")

// 	var datas = []gin.H{
// 		{"ID":      91212,
// 		"Name":     "iNew Basic Tee",
// 		"imageURL": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
// 		"price":    "$35",},
// 		{"ID":       2,
// 		"Name":     "Basic Tee",
// 		"imageURL": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
// 		"price":    "$135"},
// }
	// log.Println("Ping endpoint was hit")
	if categoryNames == "" {
		datas := model.GetAll()
		c.JSON(200, gin.H{"datas": datas})
	}else{
		categoryList := strings.Split(categoryNames, ",")
		datas, err := model.GetByCategories(categoryList)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Categories not found"})
			return
		}
		c.JSON(200, gin.H{"datas": datas})
	}
}

