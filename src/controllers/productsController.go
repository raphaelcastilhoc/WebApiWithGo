package controllers

import (
	"net/http"

	"WebApiWithGo/datastore"
	"WebApiWithGo/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := datastore.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else if len(products) > 0 {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
	}
}

func GetProductById(c *gin.Context) {
	product, err := datastore.GetProductById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else if product != nil {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
	}
}

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := datastore.AddProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	err := datastore.UpdateProduct(c.Param("id"), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}

func DeleteProduct(c *gin.Context) {
	err := datastore.DeleteProduct(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}
