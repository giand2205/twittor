package db

import (
	"context"
	"github.com/giand2205/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{
		"$match": bson.M{
			"userID": ID,
		},
	})

	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationID",
			"foreignField": "userID",
			"as":           "tweet",
		}})

	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
