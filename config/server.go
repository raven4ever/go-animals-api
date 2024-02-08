package config

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port int `env:"SERVER_PORT,default=8585"`
}

func (sc *ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", sc.Port)
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
