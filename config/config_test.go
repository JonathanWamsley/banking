package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this test must be at the root of the dir /banking to access the .env file
func TestNewConfig(t *testing.T) {
	config := NewConfig()
	assert.NotNil(t, config.MySQL.Username)
	assert.NotNil(t, config.MySQL.Password)
	assert.NotNil(t, config.MySQL.Host)
	assert.NotNil(t, config.MySQL.Schema)
	assert.NotNil(t, config.Server.Address)
	assert.NotNil(t, config.Server.Port)
}

func TestGetMySQLInfoNoError(t *testing.T) {
	config := NewConfig()
	dbInfo := config.GetMySQLInfo()
	assert.NotNil(t, dbInfo)
}

func TestGetServerInfoNoError(t *testing.T) {
	config := NewConfig()
	dbInfo := config.GetServerInfo()
	assert.NotNil(t, dbInfo)
}
