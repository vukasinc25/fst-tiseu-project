package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Firm struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id"`
	Name                string `bson:"name" json:"name"`
	TypeOfActivity      string `bson:"typeOfActivity" json:"typeOfActivity"`
	PasswordForActivity string `bson:"passwordForActivity" json:"passwordForActivity"`
}
