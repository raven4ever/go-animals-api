package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ServerConfig
	Neo4j  Neo4jConfig
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
