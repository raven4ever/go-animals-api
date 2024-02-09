package model

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Animal struct {
	ID      string
	Name    string
	Species string
	Age     int64
}

type Animals []Animal

func CreateAnimalFromNeo4jNode(node neo4j.Node) Animal {
	return Animal{
		ID:      node.Props["ID"].(string),
		Name:    node.Props["Name"].(string),
		Species: node.Props["Species"].(string),
		Age:     node.Props["Age"].(int64),
	}
}
