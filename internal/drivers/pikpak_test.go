package drivers_test

import (
	"os"
	"testing"

	"github.com/gizmo-ds/panshare/internal/drivers"
	"github.com/stretchr/testify/assert"
)

func TestPikPak_FileList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	d, err := drivers.NewPikPak(
		os.Getenv("PIKPAK_USERNAME"),
		os.Getenv("PIKPAK_PASSWORD"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	list, err := d.FileList(os.Getenv("PIKPAK_PATH"))
	assert.NoError(t, err)
	assert.NotEmpty(t, list)

	for _, f := range list {
		t.Logf("%+v", f)
	}
}

func TestPikPak_Link(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	d, err := drivers.NewPikPak(
		os.Getenv("PIKPAK_USERNAME"),
		os.Getenv("PIKPAK_PASSWORD"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	f, err := d.Link(os.Getenv("PIKPAK_FILE"))
	assert.NoError(t, err)
	assert.NotNil(t, f)
	t.Logf("%+v", f)
}
