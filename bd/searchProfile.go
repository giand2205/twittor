package bd

import (
	"context"
	"github.com/giand2205/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Println("Record not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
