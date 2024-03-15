package usermanagementapi_test

import (
	"testing"

	usermanagementapi "github.com/Mitrajeet-Golsangi/diet-application-backend/internal/app/user-management-api"
	"github.com/stretchr/testify/assert"
)

func TestPingGet(t *testing.T) {
	usermanagementapi.PingGet()

	assert.Equal(t, "pong", "pong")
}
