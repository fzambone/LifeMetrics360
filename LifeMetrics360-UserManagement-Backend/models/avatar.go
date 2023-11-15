package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Avatar struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserEmail string             `bson:"userEmail"`
	AvatarURL string             `bson:"avatarUrl"`
}
