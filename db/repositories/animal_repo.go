package repositories

import (
	"animalz/db/model"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	createAnimalQuery = "CREATE (a:Animal {ID: $id, Name: $name, Species: $species, Age: $age}) RETURN a"
	getAnimalsQuery   = "MATCH (a:Animal) RETURN a"
	getAnimalById     = "MATCH (a:Animal {ID: $id}) RETURN a"
	deleteAnimalById  = "MATCH (a:Animal {ID: $id}) DELETE a"
)

type AnimalRepo struct {
	session neo4j.SessionWithContext
	ctx     context.Context
}

func NewAnimalRepo(session neo4j.SessionWithContext, ctx context.Context) *AnimalRepo {
	return &AnimalRepo{session: session, ctx: ctx}
}

func (r *AnimalRepo) CreateAnimal(animal model.Animal) (model.Animal, error) {
	newAnimal, err := r.session.ExecuteWrite(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx,
			createAnimalQuery,
			map[string]interface{}{
				"id":      animal.ID,
				"name":    animal.Name,
				"species": animal.Species,
				"age":     animal.Age,
			})
		if err != nil {
			return nil, err
		}

		if result.Next(r.ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if err != nil {
		fmt.Println(err)
	}

	return model.CreateAnimalFromNeo4jNode(newAnimal.(neo4j.Node)), err
}

func (r *AnimalRepo) GetAnimals() (model.Animals, error) {
	animals, err := r.session.ExecuteRead(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx, getAnimalsQuery, nil)
		if err != nil {
			return nil, err
		}
		var animals model.Animals
		for result.Next(r.ctx) {
			animals = append(animals, model.CreateAnimalFromNeo4jNode(result.Record().Values[0].(neo4j.Node)))
		}
		return animals, nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return animals.(model.Animals), err
}

func (r *AnimalRepo) GetAnimal(id string) (model.Animal, error) {
	animal, err := r.session.ExecuteRead(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx, getAnimalById, map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		if result.Next(r.ctx) {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})

	if err != nil {
		fmt.Println(err)
	}

	if animal == nil {
		return model.Animal{}, nil
	}

	return model.CreateAnimalFromNeo4jNode(animal.(neo4j.Node)), err
}

func (r *AnimalRepo) DeleteAnimal(id string) (bool, error) {
	_, err := r.session.ExecuteWrite(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx, deleteAnimalById, map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})

	if err != nil {
		fmt.Println(err)
	}

	return true, err
}
