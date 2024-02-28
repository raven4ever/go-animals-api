package db

import (
	"animalz/config"
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func New(dbConfig config.Neo4jConfig) (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(dbConfig.Addr(), neo4j.BasicAuth(dbConfig.User, dbConfig.Password, ""))
}

func NewSession(ctx context.Context, driver neo4j.DriverWithContext, accessMode neo4j.AccessMode) neo4j.SessionWithContext {
	return driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: accessMode})
}
