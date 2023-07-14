package models

import (
	"bar/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Refill struct {
		PrivateID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.Refill
	}
)

// ToJSON converts the model to JSON
func (o *Refill) ToJSON() []byte {
	data, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return data
}
