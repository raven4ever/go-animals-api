package db

import (
	"animalz/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func New(dbConfig config.Neo4jConfig) (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(dbConfig.Addr(), neo4j.BasicAuth(dbConfig.User, dbConfig.Password, ""))
}
