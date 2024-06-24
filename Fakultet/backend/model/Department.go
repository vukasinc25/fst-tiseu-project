package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID    primitive.ObjectID `json:"_id"`
	Name  string             `json:"name"`
	Staff []*User            `json:"staff"`
}

type DepartmentDB struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name" json:"name"`
}

type Departments []*DepartmentDB
