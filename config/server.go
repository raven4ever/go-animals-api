package config

import "fmt"

type ServerConfig struct {
	Port int `env:"SERVER_PORT,default=8585"`
}

func (sc *ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", sc.Port)
}
