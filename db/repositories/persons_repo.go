package repositories

import (
	"animalz/db/model"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	createPersonQuery = "CREATE (p:Person {ID: $id, FirstName: $firstName, LastName: $lastName, NickName: $nickName, Age: $age}) RETURN p"
	getPersonsQuery   = "MATCH (p:Person) RETURN p"
)

type PersonRepo struct {
	session neo4j.SessionWithContext
	ctx     context.Context
}

func NewPersonRepo(session neo4j.SessionWithContext, ctx context.Context) *PersonRepo {
	return &PersonRepo{session: session, ctx: ctx}
}

func (r *PersonRepo) CreatePerson(person model.Person) (model.Person, error) {
	newPerson, err := r.session.ExecuteWrite(r.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(r.ctx,
			createPersonQuery,
			map[string]interface{}{
				"id":        person.ID,
				"firstName": person.FirstName,
				"lastName":  person.LastName,
				"nickName":  person.NickName,
				"age":       person.Age,
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

	return model.CreatePersonFromNeo4jNode(newPerson.(neo4j.Node)), err
}
