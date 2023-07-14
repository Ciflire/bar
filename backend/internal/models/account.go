package models

import (
	"bar/autogen"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Account struct {
		PrivateID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
		autogen.Account
	}
)

// ToJSON converts the model to JSON
func (o *Account) ToJSON() []byte {
	data, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return data
}
