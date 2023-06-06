package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"WebApiWithGo/datastore"
	"WebApiWithGo/models"

	"github.com/gin-gonic/gin"
)

var productCollectionName = "products"

func GetProducts(c *gin.Context) {
	var productsCollection = datastore.MongoClient.Database(datastore.DatabaseName).Collection(productCollectionName)

	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cur, err := productsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var product models.Product
		err := cur.Decode(&product)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, products)
}
