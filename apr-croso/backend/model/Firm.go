package model

type Firm struct {
	ID                  string `bson:"id" json:"id"`
	Name                string `bson:"name" json:"name"`
	TypeOfActivity      string `bson:"typeOfActivity" json:"typeOfActivity"`
	PasswordForActivity string `bson:"passwordForActivity" json:"passwordForActivity"`
}
