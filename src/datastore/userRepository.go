package datastore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"WebApiWithGo/models"
)

var userCollectionName = "users"

func GetUsers() ([]models.User, error) {
	var userCollection = MongoClient.Database(DatabaseName).Collection(userCollectionName)

	var users []models.User

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	cur, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
