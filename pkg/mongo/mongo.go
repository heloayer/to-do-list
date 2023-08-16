package mongo

import (
	"context"
	"fmt"

	"github.com/heloayer/todo-list/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	database   *mongo.Database
	Collection *mongo.Collection
}

// Установка конфига и подключение
func New(cfg config.Mongo) (*Mongo, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.URL))
	if err != nil {
		return nil, fmt.Errorf("mongo - func New - mongo.Connect: %w", err)
	}

	db := client.Database("todo-list")
	collection := db.Collection("tasks")

	// Проверяем, что есть подключение
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("mongo - func New - client.Ping: %w", err)
	}

	fmt.Println("connected to MongoDB!")

	return &Mongo{Client: client, database: db, Collection: collection}, nil
}
