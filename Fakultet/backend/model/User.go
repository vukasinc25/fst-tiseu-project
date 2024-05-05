package model

type Role string

const (
	STUDENT   Role = "STUDENT"
	PROFESSOR Role = "PROFESSOR"
	ADMIN     Role = "ADMIN"
)

type User struct {
	ID        string `bson:"id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	JMBG      int    `bson:"jmbg" json:"jmbg"`
	UserName  string `bson:"userName" json:"userName"`
	Password  string `bson:"password" json:"password"`
	Role      Role   `bson:"role" json:"role"`
}
