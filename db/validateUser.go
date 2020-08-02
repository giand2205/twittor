package db

import (
	"context"
	"github.com/giand2205/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func ValidateUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		log.Println(err.Error())
		return result, false, ID
	}

	return result, true, ID
}
