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

type ZaposljavanjeRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*ZaposljavanjeRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}
	return &ZaposljavanjeRepository{
		cli: client, logger: logger,
	}, nil
}

func (nr *ZaposljavanjeRepository) Disconnect(ctx context.Context) error {
	err := nr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (nr *ZaposljavanjeRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := nr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		nr.logger.Println(err)
	}

	// Print available databases
	databases, err := nr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		nr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (nr *ZaposljavanjeRepository) GetAll(ctx context.Context) (model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notificationCollection := nr.getCollection()

	var users model.Users
	notificationCursor, err := notificationCollection.Find(ctx, bson.M{})
	if err != nil {
		nr.logger.Println(err)
		return nil, err
	}
	if err = notificationCursor.All(ctx, &users); err != nil {
		nr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (nr *ZaposljavanjeRepository) GetAllByHostId(id string, ctx context.Context) (model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notificationCollection := nr.getCollection()

	var users model.Users
	//objID, _ := primitive.ObjectIDFromHex(id)
	notificationCursor, err := notificationCollection.Find(ctx, bson.D{{"hostId", id}})
	if err != nil {
		nr.logger.Println(err)
		return nil, err
	}
	if err = notificationCursor.All(ctx, &users); err != nil {
		nr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (nr *ZaposljavanjeRepository) Delete(username string, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := nr.getCollection()

	// objID, _ := primitive.ObjectIDFromHex(username)
	filter := bson.M{"username": username}
	result, err := patientsCollection.DeleteMany(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (nr *ZaposljavanjeRepository) GetById(id string, ctx context.Context) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notificationCollection := nr.getCollection()

	var user model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := notificationCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		nr.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (nr *ZaposljavanjeRepository) Insert(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationCollection := nr.getCollection()
	//user.Date = time.Now()
	result, err := accommodationCollection.InsertOne(ctx, &user)
	if err != nil {
		nr.logger.Println(err)
		return err
	}
	nr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (nr *ZaposljavanjeRepository) getCollection() *mongo.Collection {
	patientDatabase := nr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("notifications")
	return patientsCollection
}
