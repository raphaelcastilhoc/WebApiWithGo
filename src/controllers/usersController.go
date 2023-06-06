package controllers

import (
	"net/http"

	"WebApiWithGo/datastore"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users = datastore.GetUsers()

	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": "DATA_NOT_FOUND", "message": "Data not found"})
	}
}
