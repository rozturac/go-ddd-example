package events

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserCreated struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	UserName  string             `json:"user_name"`
}
