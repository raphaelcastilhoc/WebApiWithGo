package main

import (
	"github.com/gin-gonic/gin"

	"WebApiWithGo/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.GET("/products", controllers.GetProducts)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run()
}
