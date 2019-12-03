package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrainingMember interface {
	AllTrainers() ([]*Trainer, error)
	InsertTrainer(Trainer) (Trainer, error)
}

type DB struct {
	*mongo.Client
}

func NewDB(dataSourceName string) (*DB, error) {
	clientOptions := options.Client().ApplyURI(dataSourceName)

	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
