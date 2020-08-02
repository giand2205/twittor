package models

type Relation struct {
	UserID string `bson:"userID" json:"userID"`
	UserRelationID string `bson:"userRelationID" json:"userRelationID"`
}