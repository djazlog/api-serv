package env

import (
	"fmt"
	"net"
	"os"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

type swaggerConfig struct {
	host string
	port string
}

func (cfg swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s is not set", swaggerHostEnvName)
	}
	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s is not set", swaggerPortEnvName)
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}
