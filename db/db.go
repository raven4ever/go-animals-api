package db

import (
	"animalz/config"
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Database struct {
	Context context.Context
	Driver  neo4j.DriverWithContext
}

func NewDatabase(dbConfig config.Neo4jConfig) (*Database, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(dbConfig.Addr(), neo4j.BasicAuth(dbConfig.User, dbConfig.Password, ""))
	if err != nil {
		return nil, err
	}

	return &Database{
		Context: ctx,
		Driver:  driver,
	}, nil
}

func (d *Database) Close() {
	d.Driver.Close(d.Context)
}

func (d *Database) VerifyConnectivity() error {
	return d.Driver.VerifyConnectivity(d.Context)
}
