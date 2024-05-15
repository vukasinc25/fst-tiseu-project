package model

type Role string

const (
	VLASNIK Role = "VLASNIK"
	ADMIN   Role = "ADMIN"
)

type user struct {
	ID        string `bson:"id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	JMBG      int    `bson:"jmbg" json:"jmbg"`
	Email     string `bson:"email" json:"email"`
	UserName  string `bson:"userName" json:"userName"`
	Password  string `bson:"password" json:"password"`
	Role      Role   `bson:"role" json:"role"`
}
