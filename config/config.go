package config

import "os"

// MySQLConfig holds env variables
type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Schema   string
}

// Config holds the MySQL Config that can be called from other files
type Config struct {
	MySQL MySQLConfig
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
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
