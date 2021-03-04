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
}
