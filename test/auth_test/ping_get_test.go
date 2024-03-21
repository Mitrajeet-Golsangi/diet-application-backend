package auth_test

import (
	"testing"

	"github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/auth"
	"github.com/stretchr/testify/assert"
)

func TestPingGet(t *testing.T) {
	auth.PingGet()

	assert.Equal(t, "pong", "pong")
}
