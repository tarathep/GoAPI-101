package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func (db *DB) AllTrainers() ([]*Trainer, error) {

	collection := db.Database("test").Collection("trainers")

	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	//fmt.Printf("Found multiple documents : %+v\n", results[0])

	return results, nil
}

func (db *DB) InsertTrainer(trainer Trainer) (Trainer, error) {
	collection := db.Database("test").Collection("trainers")
	_, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		return trainer, (err)
	}

	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return trainer, nil
}
