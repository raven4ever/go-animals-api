package config

import "fmt"

type Neo4jConfig struct {
	Protocol string `env:"NEO4J_PROTOCOL,default=bolt"`
	Host     string `env:"NEO4J_HOST,default=127.0.0.1"`
	Port     int    `env:"NEO4J_PORT,default=7687"`
	User     string `env:"NEO4J_USER,default=neo4j"`
	Password string `env:"NEO4J_PASSWORD,default=neo4jp@ssw0rd"`
	DemoData bool   `env:"NEO4J_DEMO_DATA,default=true"`
}

func (nc *Neo4jConfig) Addr() string {
	return fmt.Sprintf("%s://%s:%d", nc.Protocol, nc.Host, nc.Port)
}
