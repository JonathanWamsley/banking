package config

import (
	"fmt"
	"os"
)

// MySQLConfig holds env variables
type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Schema   string
}

// ServerConfig holds env to run the server
type ServerConfig struct {
	Address string
	Port    string
}

// Config holds the MySQL Config that can be called from other files
type Config struct {
	MySQL  MySQLConfig
	Server ServerConfig
}

// NewConfig returns a new config that looks at a .env for environment variables
func NewConfig() *Config {
	return &Config{
		MySQL: MySQLConfig{
			Username: getEnv("mysql_users_username", "user"),
			Password: getEnv("mysql_users_password", "password"),
			Host:     getEnv("mysql_users_host", "127.0.0.1:3306"),
			Schema:   getEnv("mysql_users_schema", "banking"),
		},
		Server: ServerConfig{
			Address: getEnv("server_Address", "localhost"),
			Port:    getEnv("server_port", "8000"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetMySQLInfo returns string to connect to mysql db
func (c Config) GetMySQLInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		c.MySQL.Username,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Schema,
	)
}

// GetServerInfo returns string to run sever on address:port
func (c Config) GetServerInfo() string {
	return fmt.Sprintf("%s:%s", c.Server.Address, c.Server.Port)
}
