package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vukasinc25/fst-tiseu-project/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (nr *NewRepository) Insert(newUser *model.User, ctx context.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection, err := nr.getCollection(3)
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

func (nr *NewRepository) CreateCompetition(competition *model.Competition) error {
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

func (nr *NewRepository) getCollection(id int) (*mongo.Collection, error) {
	competitionDatabase := nr.cli.Database("mongoDemo")
	var competitionCollection *mongo.Collection
	switch id {
	case 1:
		competitionCollection = competitionDatabase.Collection("competition")
	case 2:
		competitionCollection = competitionDatabase.Collection("registeredStudentsToCommpetition")
	case 3:
		competitionCollection = competitionDatabase.Collection("fakultetUsers")
	case 4:
		competitionCollection = competitionDatabase.Collection("diplomas")
	default:
		return nil, fmt.Errorf("invalid collection id")
	}

	return competitionCollection, nil

	return competitionCollection, nil
}
