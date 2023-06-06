package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"WebApiWithGo/models"
)

var productCollectionName = "products"

func GetProducts() []models.Product {
	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	var products []models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cur, err := productCollection.Find(ctx, bson.D{})
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

	return products
}
