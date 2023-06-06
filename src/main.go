package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"WebApiWithGo/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductById)
	router.POST("/products", controllers.AddProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	router.GET("/users", controllers.GetUsers)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
	})

	router.Run()
}
