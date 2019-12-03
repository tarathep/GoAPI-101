package main

import (
	"fmt"
	"log"

	"github.com/tarathep/GoAPI-101/trainer/models"
)

type Env struct {
	db models.TrainingMember
}

func main() {
	db, err := models.NewDB("mongodb://admin:password@localhost:27017")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	env.trainersFindAll()
}

func (env *Env) trainersFindAll() {
	ts, err := env.db.AllTrainers()

	if err != nil {
		println("error 500 !!")
		return
	}

	for _, t := range ts {
		fmt.Println(t.Name, t.City, t.Age)
	}
}

func (env *Env) trainerInsertOne(trainer models.Trainer) {
	result, err := env.db.InsertTrainer(trainer)
	fmt.Println("Inserted a single document: ", result.Name, err)
}
