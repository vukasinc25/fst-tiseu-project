package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobListing struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	JobTitle       string             `bson:"hostId, omitempty" json:"jobTitle"`
	JobDescription string             `bson:"description,omitempty" json:"jobDescription"`
	Requirements   string             `bson:"date,omitempty" json:"requirements"`
}

type JobListings []*JobListing
