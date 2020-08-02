package db

import (
	"context"
	"github.com/giand2205/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	condition := bson.M{
		"userID":         t.UserID,
		"userRelationID": t.UserRelationID,
	}

	var result models.Relation
	log.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
