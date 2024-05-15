package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competition struct {
	ID                   primitive.ObjectID
	ProgramName          string
	AdmissionRequrements string
	ExamDate             time.Time
	ExamFormat           string
	ExamMaterials        string
	ApplicationDeadlines string
	ApplicationDocuments string
	ApplicationMethod    string
	ApplicationContact   string
	TutionFees           string
	ContactInformation   string
}
