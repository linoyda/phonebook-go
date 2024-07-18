package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    FirstName string             `json:"first_name" bson:"first_name"`
    LastName  string             `json:"last_name" bson:"last_name"`
    Phone     string             `json:"phone" bson:"phone"`
    Address   string             `json:"address" bson:"address"`
}
