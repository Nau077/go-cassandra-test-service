package config

import (
	"encoding/json"
	"net"
	"os"
)

const (
	dbPassEscSeq = "{password}"
	password     = "note-service-password"
)

// HTTP
type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}

// db
type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

// config
type Config struct {
	HTTP HTTP `json:"http"`
	DB   DB   `json:"db"`
}

// new config
func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
