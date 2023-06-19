package bark_test

import (
	"os"
	"testing"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/bark"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func LoadEnv() error {
	return godotenv.Load("../../.env")
}

func TestPush(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	err = bark.Push(os.Getenv("BARK_KEY"), utils.M{
		"body":  "Test Bark Server",
		"title": "Test Title",
		"icon":  "https://day.app/assets/images/avatar.jpg",
	})
	assert.NoError(t, err)
}
