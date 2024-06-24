package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vukasinc25/fst-tiseu-project/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type NewRepository struct {
	cli *mongo.Client
}

func New(ctx context.Context) (*NewRepository, error) {
	dbURI := "mongodb://root:pass@mongo:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	return &NewRepository{
		cli: client,
	}, nil
}

func (uh *NewRepository) Disconnect(ctx context.Context) error {
	err := uh.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (nr *NewRepository) InsertDepartment(department *model.DepartmentDB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	departmentCollection, err := nr.getCollection(6)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := departmentCollection.InsertOne(ctx, department)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (uh *NewRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := uh.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Ping error 1: ", err)
	}

	databases, err := uh.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println("Ping error 2: ", err)
	}
	fmt.Println(databases)
}

func (nr *NewRepository) GetAllCompetitions() (*model.Competitions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection, err := nr.getCollection(1)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return nil, err
	}

	var competitons model.Competitions
	cursor, err := competitionCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Cant find departmentCollection: ", err)
		return nil, err
	}
	if err = cursor.All(ctx, &competitons); err != nil {
		log.Println("Department Cursor.All: ", err)
		return nil, err
	}
	return &competitons, nil
}

func (nr *NewRepository) Insert(newUser *model.User, ctx context.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection, err := nr.getCollection(7)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := usersCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) InsertStudyProgram(studyProgram *model.StudyProgram) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	studyProgramCollection, err := nr.getCollection(8)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := studyProgramCollection.InsertOne(ctx, studyProgram)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) InsertCompetition(competition *model.Competition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection, err := nr.getCollection(1)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := competitionCollection.InsertOne(ctx, competition)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) GetAllStudyPrograms() (*model.StudyPrograms, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	studyProgramCollection, err := nr.getCollection(8)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return nil, err
	}

	var studyPrograms model.StudyPrograms
	cursor, err := studyProgramCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Cant find departmentCollection: ", err)
		return nil, err
	}
	if err = cursor.All(ctx, &studyPrograms); err != nil {
		log.Println("Department Cursor.All: ", err)
		return nil, err
	}
	return &studyPrograms, nil
}

func (nr *NewRepository) CreateRegisteredStudentToTheCommpetition(registeredStudentsToCommpetition *model.RegisteredStudentsToCommpetition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registeredStudentsToCommpetitionCollection, err := nr.getCollection(2)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := registeredStudentsToCommpetitionCollection.InsertOne(ctx, registeredStudentsToCommpetition)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) GetAllDepartments() (*model.Departments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	departmentCollection, err := nr.getCollection(6)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return nil, err
	}

	var departments model.Departments
	cursor, err := departmentCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Cant find departmentCollection: ", err)
		return nil, err
	}
	if err = cursor.All(ctx, &departments); err != nil {
		log.Println("Department Cursor.All: ", err)
		return nil, err
	}
	return &departments, nil
}

func (nr *NewRepository) GetAllUsers() (*model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userCollection, err := nr.getCollection(7)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return nil, err
	}

	var users model.Users
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Cant find departmentCollection: ", err)
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		log.Println("Department Cursor.All: ", err)
		return nil, err
	}
	return &users, nil
}

func (nr *NewRepository) InsertUserExamResult(results *model.ExamResult) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection, err := nr.getCollection(5)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := competitionCollection.InsertOne(ctx, results)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) GetAllExamResultsByCompetitionId(competitionId string) (*model.ExamResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// competitionId = "6658d76eed49f71587b7c4b1" // Note: This line is for testing purposes and can be removed

	resultatCollection, err := nr.getCollection(5)
	if err != nil {
		log.Println("Error getting collection: ", err)
		return nil, err
	}

	var examResults model.ExamResults

	cursor, err := resultatCollection.Find(ctx, bson.M{"competitionId": competitionId})
	if err != nil {
		log.Println("Error finding exam results: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	log.Println("Pre petlje")

	for cursor.Next(ctx) {
		log.Println("Petlja")
		var result model.ExamResult
		if err := cursor.Decode(&result); err != nil {
			log.Println("Error decoding exam result:", err)
			return nil, err
		}
		log.Println("Decoded result:", result)
		examResults = append(examResults, &result)
	}

	log.Println("Posle petlje")
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	log.Println("Results: ", examResults)

	return &examResults, nil
}

func (nr *NewRepository) GetStudyProgramId(id string) (*model.StudyProgram, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	studyProgramCollection, err := nr.getCollection(8)
	if err != nil {
		log.Println("Error getting collection: ", err)
		return nil, err
	}

	var studyProgram model.StudyProgram

	objId, _ := primitive.ObjectIDFromHex(id)
	err = studyProgramCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&studyProgram)
	if err != nil {
		log.Println("Error decoding user document: ", err)
		return nil, err
	}
	return &studyProgram, nil
}

func (nr *NewRepository) InsertDiploma(diploma *model.Diploma) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	diplomaCollection, err := nr.getCollection(4)
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := diplomaCollection.InsertOne(ctx, diploma)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) GetDiplomaByUserId(userId string) (*model.Diploma, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	diplomaCollection, err := nr.getCollection(4)
	if err != nil {
		log.Println("Error getting collection: ", err)
		return nil, err
	}
	var diploma model.Diploma
	log.Println("Querying for user with id: ", userId)

	err = diplomaCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&diploma)
	if err != nil {
		log.Println("Error decoding user document: ", err)
		return nil, err
	}
	return &diploma, nil
}

func (nr *NewRepository) GetCompetitionById(id string) (*model.Competition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection, err := nr.getCollection(1)
	if err != nil {
		log.Println("Error getting collection: ", err)
		return nil, err
	}
	var competition model.Competition
	objId, _ := primitive.ObjectIDFromHex(id)
	err = competitionCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&competition)
	if err != nil {
		log.Println("Error decoding user document: ", err)
		return nil, err
	}
	return &competition, nil
}

func (nr *NewRepository) getCollection(id int) (*mongo.Collection, error) {
	competitionDatabase := nr.cli.Database("mongoDemo")
	var competitionCollection *mongo.Collection
	switch id {
	case 1:
		competitionCollection = competitionDatabase.Collection("competitions")
	case 2:
		competitionCollection = competitionDatabase.Collection("registeredStudentsToCommpetition")
	// case 3:
	// 	competitionCollection = competitionDatabase.Collection("fakultetUsers")
	case 4:
		competitionCollection = competitionDatabase.Collection("diplomas")
	case 5:
		competitionCollection = competitionDatabase.Collection("examResults")
	case 6:
		competitionCollection = competitionDatabase.Collection("departments")
	case 7:
		competitionCollection = competitionDatabase.Collection("users")
	case 8:
		competitionCollection = competitionDatabase.Collection("studyPrograms")
	default:
		return nil, fmt.Errorf("invalid collection id")
	}

	return competitionCollection, nil
}
