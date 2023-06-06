package controllers

import (
	"net/http"

	"WebApiWithGo/datastore"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products = datastore.GetProducts()

	if len(products) > 0 {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": "DATA_NOT_FOUND", "message": "Data not found"})
	}
}
