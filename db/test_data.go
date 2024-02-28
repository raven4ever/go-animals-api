package db

import (
	"animalz/db/model"
	"animalz/db/repositories"

	"log"

	"github.com/google/uuid"
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

func (d *Database) LoadDemoData() {
	session := d.Driver.NewSession(d.Context, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(d.Context)

	// load persons
	pRepo := repositories.NewPersonRepo(session, d.Context)

	for _, p := range testPersons {
		_, err := pRepo.CreatePerson(p)
		if err != nil {
			log.Println(err)
		}
	}

	// load animals
	aRepo := repositories.NewAnimalRepo(session, d.Context)

	for _, a := range testAnimals {
		_, err := aRepo.CreateAnimal(a)
		if err != nil {
			log.Println(err)
		}
	}

	// load foods
	fRepo := repositories.NewFoodRepo(session, d.Context)

	for _, f := range testFoods {
		_, err := fRepo.CreateFood(f)
		if err != nil {
			log.Println(err)
		}
	}
}
