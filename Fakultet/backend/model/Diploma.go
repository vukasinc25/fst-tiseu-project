package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diploma struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserId       string             `bson:"userId" json:"userId"`
	UserName     string             `bson:"userName" json:"userName"`
	IssueDate    time.Time          `bson:"issueDate" json:"issueDate"`
	AverageGrade string             `bson:"averageGrade" json:"averageGrade"`
}

type DiplomaRequest struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	UserId     string             `bson:"userId" json:"userId"`
	UserName   string             `bson:"userName" json:"userName"`
	IssueDate  time.Time          `bson:"issueDate" json:"issueDate"`
	InPending  bool               `bson:"inPending"`
	IsApproved bool               `bson:"isApproved"`
}

type IsApproved struct {
	IsApproved bool `json:"isApproved"`
}

type DiplomaRequests []*DiplomaRequest
