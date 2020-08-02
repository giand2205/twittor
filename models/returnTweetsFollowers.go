package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReturnTweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         string             `bson:"userID" json:"userID,omitempty"`
	UserRelationID string             `bson:"userRelationID" json:"userRelationID,omitempty"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id,omitempty" json:"id"`
	}
}
