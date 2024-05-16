package repository

import (
	"context"
	"fmt"
	"github.com/vukasinc25/fst-tiseu-project/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type ListingRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*ListingRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}
	return &ListingRepository{
		cli: client, logger: logger,
	}, nil
}

func (lr *ListingRepository) Disconnect(ctx context.Context) error {
	err := lr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (lr *ListingRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := lr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		lr.logger.Println(err)
	}

	// Print available databases
	databases, err := lr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		lr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (lr *ListingRepository) Insert(jobListing *model.JobListing) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jobListingCollection := lr.getCollection()
	//if err != nil {
	//	return errors.New("error in getting jobListing collection")
	//}

	result, err := jobListingCollection.InsertOne(ctx, &jobListing)
	if err != nil {
		log.Println("Errsor when tryed to insert jobListing: ", err)
		return err
	}
	log.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (lr *ListingRepository) GetAll(ctx context.Context) (model.JobListings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobListingCollection := lr.getCollection()

	var users model.JobListings
	jobListingCursor, err := jobListingCollection.Find(ctx, bson.M{})
	if err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	if err = jobListingCursor.All(ctx, &users); err != nil {
		lr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (lr *ListingRepository) getCollection() *mongo.Collection {
	patientDatabase := lr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("jobListings")
	return patientsCollection
}
