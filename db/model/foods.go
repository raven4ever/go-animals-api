package model

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Food struct {
	ID          string
	Name        string
	Description string
	Qty         int64
}

type Foods []Food

func CreateFoodFromNeo4jNode(node neo4j.Node) Food {
	return Food{
		ID:          node.Props["ID"].(string),
		Name:        node.Props["Name"].(string),
		Description: node.Props["Description"].(string),
		Qty:         node.Props["Qty"].(int64),
	}
}
