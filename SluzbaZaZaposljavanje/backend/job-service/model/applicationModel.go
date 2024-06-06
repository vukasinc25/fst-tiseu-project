package model

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type JobApplication struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	JobListingId string             `bson:"jobListingId, omitempty" json:"jobListingId"`
	EmployerId   string             `bson:"employerId,omitempty" json:"employerId"`
	EmployeeId   string             `bson:"employeeId,omitempty" json:"employeeId"`
	Diploma      string             `bson:"diploma,omitempty" json:"diploma"`
}

type JobApplications []*JobApplication

func (as *JobApplications) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(as)
}

func (a *JobApplications) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func (a *JobApplication) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *JobApplication) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

func ValidateJobApplication(notification *JobApplication) error {
	if notification.JobListingId == "" {
		return errors.New("job title is required")
	}
	if notification.EmployeeId == "" {
		return errors.New("job description is required")
	}

	return nil
}
