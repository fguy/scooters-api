package scooters

import (
	"testing"

	"github.com/fguy/scooters-api/config"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestNewHTTPServer(t *testing.T) {
	t.Parallel()

	server := NewHTTPServer(&config.AppConfig{}, zap.NewNop(), nil, nil)
	assert.NotNil(t, server)
}
