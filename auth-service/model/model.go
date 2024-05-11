package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqToken struct {
	Token string `json:"token"`
}

type Role string

const (
	Host  Role = "HOST"
	Guest Role = "GUEST"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username,omitempty" json:"username" validate:"required,min=6"`
	Password  string             `bson:"password,omitempty" json:"password" validate:"required,password"`
	Role      Role               `bson:"role,omitempty" json:"role" validate:"required,oneof=HOST GUEST"`
	Email     string             `bson:"email,omitempty" json:"email" validate:"required,email"`
	FirstName string             `bson:"firstname,omitempty" json:"firstname"`
	LastName  string             `bson:"lastname,omitempty" json:"lastname"`
}

type ResponseUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username"`
	Role     string             `bson:"role,omitempty" json:"role"`
}

type Users []*ResponseUser
