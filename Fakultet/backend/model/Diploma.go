package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diploma struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserId       string             `bson:"userId" json:"userId"`
	IssueDate    time.Time          `bson:"issueDate" json:"issueDate"`
	AverageGrade string             `bson:"averageGrade" json:"averageGrade"`
}
