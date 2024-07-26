package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tinaba96/taka-ultimate-fashion/backend/controller"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Href     string `json:"href"`
	ImageSrc string `json:"imageSrc"`
	ImageAlt string `json:"imageAlt"`
	Price    string `json:"price"`
	Color    string `json:"color"`
}

func SetRouter() *gin.Engine {
	r := gin.Default()


	r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:4000"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }))
	// r.GET("/products", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Health check OK",
	// 	})
	// })
// 	r.GET("/test", func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")

//         c.JSON(200, []gin.H{
//                 {"id":      100,
//                 "name":     "iNew Basic Tee",
//                 "href":     "#",
//                 "imageSrc": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
//                 "imageAlt": "Front of men's Basic Tee in black.",
//                 "price":    "$35",
//                 "color":    "White",},
//                 {"id":       2,
//                 "name":     "Basic Tee",
//                 "href":     "#",
//                 "imageSrc": "https://tailwindui.com/img/ecommerce-images/product-page-01-related-product-01.jpg",
//                 "imageAlt": "Back of men's Basic Tee in black.",
//                 "price":    "$135",
//                 "color":    "Black",},
// })
//     })
    r.GET("/products", controller.ShowAllProducts)
    r.GET("/getSave", controller.GetSave)
    r.POST("/save", controller.SaveProduct)



	return r
}
