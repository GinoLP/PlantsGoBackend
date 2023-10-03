package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	LatinName string             `bson:"latinName,omitempty"`
	Name      string             `bson:"name,omitempty"`
}
