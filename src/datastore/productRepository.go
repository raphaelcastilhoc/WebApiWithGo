package datastore

import (
	"context"
	"time"

	"WebApiWithGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productCollectionName = "products"

func GetProducts() ([]models.Product, error) {
	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	var products []models.Product

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	cur, err := productCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var product models.Product

		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(id string) (*models.Product, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	var product *models.Product
	result := productCollection.FindOne(context.TODO(), bson.M{"_id": objectId})

	decodeError := result.Decode(&product)
	if decodeError != nil {
		return nil, decodeError
	}

	return product, nil
}

func AddProduct(product *models.Product) error {
	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	product.Id = primitive.NewObjectID()
	_, err := productCollection.InsertOne(context.TODO(), &product)

	return err
}

func UpdateProduct(id string, product *models.Product) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	product.Id = objectId

	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	filter := bson.D{{"_id", objectId}}
	_, updateError := productCollection.ReplaceOne(context.TODO(), filter, &product)

	return updateError
}

func DeleteProduct(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var productCollection = MongoClient.Database(DatabaseName).Collection(productCollectionName)

	filter := bson.D{{"_id", objectId}}
	_, deleteError := productCollection.DeleteOne(context.TODO(), filter)

	return deleteError
}
