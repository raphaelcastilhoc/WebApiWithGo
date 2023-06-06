package controllers

import (
	"net/http"

	"WebApiWithGo/datastore"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := datastore.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else if len(users) > 0 {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
	}
}
