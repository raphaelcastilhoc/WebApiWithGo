package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"WebApiWithGo/models"
)

var userCollectionName = "users"

func GetUsers() []models.User {
	var userCollection = MongoClient.Database(DatabaseName).Collection(userCollectionName)

	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cur, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
