package repository

import (
	"context"
	"fmt"
	"github.com/vukasinc25/fst-tiseu-project/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type JobRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*JobRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}
	return &JobRepository{
		cli: client, logger: logger,
	}, nil
}

func (jr *JobRepository) Disconnect(ctx context.Context) error {
	err := jr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (jr *JobRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := jr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		jr.logger.Println(err)
	}

	// Print available databases
	databases, err := jr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		jr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (jr *JobRepository) InsertJobListing(jobListing *model.JobListing) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jobListingCollection := jr.getJobListingCollection()
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

func (jr *JobRepository) InsertJobApplication(jobApplication *model.JobApplication) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jobApplicationCollection := jr.getJobApplicationCollection()
	//if err != nil {
	//	return errors.New("error in getting jobListing collection")
	//}

	result, err := jobApplicationCollection.InsertOne(ctx, &jobApplication)
	if err != nil {
		log.Println("Error when tried to insert jobApplication: ", err)
		return err
	}
	log.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (jr *JobRepository) GetAllJobListings() (model.JobListings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobListingCollection := jr.getJobListingCollection()

	var jobListings model.JobListings
	jobListingCursor, err := jobListingCollection.Find(ctx, bson.M{})
	if err != nil {
		jr.logger.Println(err)
		return nil, err
	}
	if err = jobListingCursor.All(ctx, &jobListings); err != nil {
		jr.logger.Println(err)
		return nil, err
	}
	return jobListings, nil
}

func (jr *JobRepository) GetJobListing(id string) (model.JobListing, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobListingCollection := jr.getJobListingCollection()

	log.Println("Fetching job listing with ID:", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		jr.logger.Println("Invalid ID format:", err)
		return model.JobListing{}, err
	}

	var job model.JobListing
	err = jobListingCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&job)
	if err != nil {
		jr.logger.Println("Error fetching job listing:", err)
		return model.JobListing{}, err
	}
	return job, nil
}

func (jr *JobRepository) GetAllJobApplicationsByEmployerId(id string) (model.JobApplications, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobApplicationCollection := jr.getJobApplicationCollection()

	//objID, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	jr.logger.Println("Invalid ID format:", err)
	//	return model.JobApplications{}, err
	//}

	var jobApplications model.JobApplications
	jobListingCursor, err := jobApplicationCollection.Find(ctx, bson.M{"employerId": id})
	if err != nil {
		jr.logger.Println(err)
		return nil, err
	}
	if err = jobListingCursor.All(ctx, &jobApplications); err != nil {
		jr.logger.Println(err)
		return nil, err
	}

	return jobApplications, nil
}

func (jr *JobRepository) getJobListingCollection() *mongo.Collection {
	patientDatabase := jr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("jobListings")
	return patientsCollection
}

func (jr *JobRepository) getJobApplicationCollection() *mongo.Collection {
	patientDatabase := jr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("jobApplications")
	return patientsCollection
}
