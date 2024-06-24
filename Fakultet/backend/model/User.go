package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	STUDENT   Role = "STUDENT"
	PROFESSOR Role = "PROFESSOR"
	ADMIN     Role = "ADMIN"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username,omitempty" json:"username" validate:"required,min=6"`
	Password  string             `bson:"password,omitempty" json:"password" validate:"required,password"`
	Role      Role               `bson:"role,omitempty" json:"role"`
	Email     string             `bson:"email,omitempty" json:"email" validate:"required,email"`
	FirstName string             `bson:"firstname,omitempty" json:"firstname"`
	LastName  string             `bson:"lastname,omitempty" json:"lastname"`
}

type Users []*User
