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

func (nr *NewRepository) CreateFirm(firm *model.Firm) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	firmCollection, err := nr.getCollection()
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := firmCollection.InsertOne(ctx, firm)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}

func (nr *NewRepository) getCollection() (*mongo.Collection, error) {
	firmDatabase := nr.cli.Database("mongoDemo")
	firmCollection := firmDatabase.Collection("firm")

	return firmCollection, nil
}
