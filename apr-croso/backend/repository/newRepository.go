package repository

import (
	"context"
	"log"
	"time"

	"github.com/vukasinc25/fst-tiseu-project/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewRepository struct {
	cli *mongo.Client
}

func New(ctx context.Context) (*NewRepository, error) {
	// dbURI := os.Getenv("MONGO_DB_URI")
	dbURI := "mongodb://root:pass@mongo:27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	// return &NewRepository{
	// 	cli: client,
	// }, nil
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
func (nr *NewRepository) CreateFirm(firm *model.Firm) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	firmsCollection, err := nr.getCollection()
	if err != nil {
		log.Println("Duplicate key error: ", err)
		return err
	}

	result, err := firmsCollection.InsertOne(ctx, firm)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Document ID: %v\n", result.InsertedID)

	return nil
}


func (nr *NewRepository) getCollection() (*mongo.Collection, error) {
	firmsDatabase := nr.cli.Database("mongoDemo")
	firmsCollection := firmsDatabase.Collection("firms")
	return firmsCollection, nil
}




