package config

import "fmt"

type Neo4jConfig struct {
	Protocol string
	Host     string
	Port     int
	User     string
	Password string
	DemoData bool
}

func (nc *Neo4jConfig) Addr() string {
	return fmt.Sprintf("%s://%s:%d", nc.Protocol, nc.Host, nc.Port)
}
