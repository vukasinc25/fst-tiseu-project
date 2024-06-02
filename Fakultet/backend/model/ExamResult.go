package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExamResult struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	StudentUserName string             `bson:"userName" json:"userName"`
	CompetitionID   string             `bson:"competitionId" json:"competitionId"`
	Score           string             `bson:"score" json:"score"`
	ScoreEntryDate  time.Time          `bson:"scoreEntryDate" json:"scoreEntryDate"`
}

type ExamResults []*ExamResult
