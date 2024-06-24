package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Competition struct {
	ID                    primitive.ObjectID `bson:"_id" json:"_id"`
	ProgramName           string             `bson:"programName" json:"programName"`
	AdmissionRequirements string             `bson:"admissionRequirements" json:"admissionRequirements"`
	ExamDate              time.Time          `bson:"examDate" json:"examDate"`
	ExamFormat            string             `bson:"examFormat" json:"examFormat"`
	ExamMaterials         string             `bson:"examMaterials" json:"examMaterials"`
	ApplicationDeadlines  string             `bson:"applicationDeadlines" json:"applicationDeadlines"`
	ApplicationDocuments  string             `bson:"applicationDocuments" json:"applicationDocuments"`
	ApplicationMethod     string             `bson:"applicationMethod" json:"applicationMethod"`
	ApplicationContact    string             `bson:"applicationContact" json:"applicationContact"`
	TuitionFees           string             `bson:"tuitionFees" json:"tuitionFees"`
	ContactInformation    string             `bson:"contactInformation" json:"contactInformation"`
}

type RegisteredStudentsToCommpetition struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	CompetitionID string             `bson:"competitionID" json:"competitionID"`
	UserID        string             `bson:"userID" json:"userID"`
}
