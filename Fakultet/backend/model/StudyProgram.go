package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudyProgram struct {
	ID                     primitive.ObjectID `bson:"_id" json:"_id"`
	Name                   string             `bson:"name" json:"name"`
	StudyLevel             string             `bson:"studyLevel" json:"studyLevel"`
	Duration               string             `bson:"duration" json:"duration"`
	Objectives             string             `bson:"objectives" json:"objectives"`
	ProgramStructure       string             `bson:"programStructure" json:"programStructure"`
	Internship             bool               `bson:"internship" json:"internship"`
	GraduationRequirements string             `bson:"graduationRequirements" json:"graduationRequirements"`
	Accreditation          bool               `bson:"accreditation" json:"accreditation"`
	ContactPersonID        string             `bson:"contactPersonID" json:"contactPersonID"`
	DevelopmentPlan        string             `bson:"developmentPlan" json:"developmentPlan"`
	DepartmentID           string             `bson:"departmentID" json:"departmentID"`
}

type StudyPrograms []*StudyProgram
