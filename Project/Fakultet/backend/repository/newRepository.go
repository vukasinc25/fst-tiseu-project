package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type NewRepository struct {
	cli *mongo.Client
}

func New(ctx context.Context) (*NewRepository, error) {
	// dbURI := os.Getenv("MONGO_DB_URI")

	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	// if err != nil {
	// 	return nil, err
	// }

	// return &NewRepository{
	// 	cli: client,
	// }, nil

	return nil, nil
}
