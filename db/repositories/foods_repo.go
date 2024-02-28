package repositories

import (
	"animalz/db/model"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	createFoodQuery = "CREATE (f:Food {ID: $id, Name: $name, Description: $description, Qty: $qty}) RETURN f"
	getFoodsQuery   = "MATCH (f:Food) RETURN f"
)

type FoodRepo struct {
	session neo4j.SessionWithContext
	ctx     context.Context
}

func NewFoodRepo(session neo4j.SessionWithContext, ctx context.Context) *FoodRepo {
	return &FoodRepo{session: session, ctx: ctx}
}

func (r *FoodRepo) CreateFood(food model.Food) (model.Food, error) {
	newFood, err := r.session.ExecuteWrite(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx,
			createFoodQuery,
			map[string]interface{}{
				"id":          food.ID,
				"name":        food.Name,
				"description": food.Description,
				"qty":         food.Qty,
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

	return model.CreateFoodFromNeo4jNode(newFood.(neo4j.Node)), err
}

func (r *FoodRepo) GetFoods() (model.Foods, error) {
	foods, err := r.session.ExecuteRead(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx, getFoodsQuery, nil)
		if err != nil {
			return nil, err
		}

		var foods model.Foods
		for result.Next(r.ctx) {
			foods = append(foods, model.CreateFoodFromNeo4jNode(result.Record().Values[0].(neo4j.Node)))
		}

		return foods, nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return foods.(model.Foods), err
}
