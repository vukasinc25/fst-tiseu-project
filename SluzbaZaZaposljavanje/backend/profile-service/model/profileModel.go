package model

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"hostId, omitempty" json:"name"`
	Address string             `bson:"description,omitempty" json:"address"`
	Email   string             `bson:"date,omitempty" json:"email"`
}

type Users []*User

type ReqToken struct {
	Token string `json:"token"`
}

func (as *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(as)
}

func (a *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func ValidateUser(notification *User) error {
	if notification.Name == "" {
		return errors.New("name is required")
	}
	if notification.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
