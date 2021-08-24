package dbrepos

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	dbClient *mongo.Client
}

var (
	Repo *Repository
)

func CreateNewDBClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func InitRepository(client *mongo.Client) {
	Repo = &Repository{dbClient: client}
}

func ConnectDBClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Repo.dbClient.Connect(ctx); err != nil {
		return err
	}

	log.Println("MongoDB Connected")
	return nil
}

func DisconnectDBClient() {
	if err := Repo.dbClient.Disconnect(context.Background()); err != nil {
		log.Println(err.Error())
	}

	log.Println("MongoDB Disconnected")
}
