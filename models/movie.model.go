package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Rating    int                `json:"rating"`
	Watched bool               `json:"watched"`
}
