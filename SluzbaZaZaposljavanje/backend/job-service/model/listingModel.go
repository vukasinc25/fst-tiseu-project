package model

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type JobListing struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	EmployerId     string             `bson:"employerId, omitempty" json:"employerId"`
	CompanyName    string             `bson:"companyName, omitempty" json:"companyName"`
	CityName       string             `bson:"cityName, omitempty" json:"cityName"`
	JobTitle       string             `bson:"jobTitle, omitempty" json:"jobTitle"`
	JobDescription string             `bson:"jobDescription,omitempty" json:"jobDescription"`
	Requirements   string             `bson:"requirements,omitempty" json:"requirements"`
}

type JobListings []*JobListing

type ReqToken struct {
	Token string `json:"token"`
}

func (as *JobListings) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(as)
}

func (a *JobListings) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func (a *JobListing) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *JobListing) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func ValidateJobListing(notification *JobListing) error {
	if notification.JobTitle == "" {
		return errors.New("job title is required")
	}
	if notification.JobDescription == "" {
		return errors.New("job description is required")
	}

	return nil
}
