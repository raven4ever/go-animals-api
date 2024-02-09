package db

import (
	"animalz/db/model"
	"animalz/db/repositories"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	testPersons = model.Persons{
		model.Person{ID: uuid.NewString(), FirstName: "John", LastName: "Doe", NickName: "JD", Age: 25},
		model.Person{ID: uuid.NewString(), FirstName: "Tracy", LastName: "Reagan", NickName: "TR", Age: 44},
		model.Person{ID: uuid.NewString(), FirstName: "Adrian", LastName: "Junior", NickName: "AJ", Age: 31},
	}

	testAnimals = model.Animals{
		model.Animal{ID: uuid.NewString(), Name: "Fido", Species: "Dog", Age: 3},
		model.Animal{ID: uuid.NewString(), Name: "Whiskers", Species: "Cat", Age: 5},
		model.Animal{ID: uuid.NewString(), Name: "Fluffy", Species: "Rabbit", Age: 2},
	}

	testFoods = model.Foods{
		model.Food{ID: uuid.NewString(), Name: "Pizza", Description: "Cheese and Tomato", Qty: 5},
		model.Food{ID: uuid.NewString(), Name: "Burger", Description: "Cheese and Beef", Qty: 10},
		model.Food{ID: uuid.NewString(), Name: "Pasta", Description: "Cheese and Tomato", Qty: 8},
	}
)

func LoadDemoData(ctx context.Context, driver neo4j.DriverWithContext) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// load persons
	pRepo := repositories.NewPersonRepo(session, ctx)

	for _, p := range testPersons {
		_, err := pRepo.CreatePerson(p)
		if err != nil {
			log.Error(err)
		}
	}

	// load animals
	aRepo := repositories.NewAnimalRepo(session, ctx)

	for _, a := range testAnimals {
		_, err := aRepo.CreateAnimal(a)
		if err != nil {
			log.Error(err)
		}
	}

	// load foods
	fRepo := repositories.NewFoodRepo(session, ctx)

	for _, f := range testFoods {
		_, err := fRepo.CreateFood(f)
		if err != nil {
			log.Error(err)
		}
	}

}
