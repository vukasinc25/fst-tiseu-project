package model

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Company struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name           string             `bson:"employerId, omitempty" json:"employerId"`
	Address        string             `bson:"jobTitle, omitempty" json:"jobTitle"`
	TotalEmployees string             `bson:"jobDescription,omitempty" json:"jobDescription"`
	Requirements   string             `bson:"requirements,omitempty" json:"requirements"`
}

type Companies []*Companies

func (as *Companies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(as)
}

func (a *Companies) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func (a *Company) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Company) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func ValidateCompanies(notification *JobListing) error {
	if notification.JobTitle == "" {
		return errors.New("job title is required")
	}
	if notification.JobDescription == "" {
		return errors.New("job description is required")
	}

	return nil
}
