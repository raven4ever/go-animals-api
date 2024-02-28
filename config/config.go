package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
	Neo4j  Neo4jConfig
}

var EnvConfig = loadEnvironmentConfig()

func loadEnvironmentConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: stringToPort(getEnvVariable("PORT", "8585")),
		},
		Neo4j: Neo4jConfig{
			Protocol: getEnvVariable("NEO4J_PROTOCOL", "bolt"),
			Host:     getEnvVariable("NEO4J_HOST", "127.0.0.1"),
			Port:     stringToPort(getEnvVariable("NEO4J_PORT", "7687")),
			User:     getEnvVariable("NEO4J_USER", "neo4j"),
			Password: getEnvVariable("NEO4J_PASSWORD", "neo4jp@ssw0rd"),
			DemoData: stringToBool(getEnvVariable("NEO4J_DEMO_DATA", "true")),
		},
	}
}

func getEnvVariable(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func stringToPort(port string) int {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("%s is not a valid port number", port)
	}
	return intPort
}

func stringToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("%s is not a valid boolean", s)
	}
	return b
}
