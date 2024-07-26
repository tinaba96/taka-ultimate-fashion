package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinaba96/taka-ultimate-fashion/backend/model"
)

type requestBodyType struct {
	ProductID uint   `json:"productID" binding:"required"`
	Status   uint 
	// Status   uint   `json:"status" binding:"required"`
}

func GetSave(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	datas := model.GetAllSaveProduct()
	c.JSON(200, gin.H{"datas": datas})
}

func SaveProduct(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	var requestBody requestBodyType

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveData := model.RequestBodyType{
        ProductID: requestBody.ProductID,
        Status:    requestBody.Status,
    }

	// err := model.PostProduct(requestBodyType{ProductID:12, Status:1})
	err := model.PostProduct(saveData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ERROR on SAVE"})
		return
	}




	c.JSON(200, gin.H{"datas": "PRODUCT SAVED!"})
}

