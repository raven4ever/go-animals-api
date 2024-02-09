package model

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Person struct {
	ID        string
	FirstName string
	LastName  string
	NickName  string
	Age       int64
}

type Persons []Person

func CreatePersonFromNeo4jNode(node neo4j.Node) Person {
	return Person{
		ID:        node.Props["ID"].(string),
		FirstName: node.Props["FirstName"].(string),
		LastName:  node.Props["LastName"].(string),
		NickName:  node.Props["NickName"].(string),
		Age:       node.Props["Age"].(int64),
	}
}
