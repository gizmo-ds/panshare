package pikpak_test

import (
	"os"
	"testing"

	"github.com/gizmo-ds/panshare/pkg/pikpak"

	"github.com/stretchr/testify/assert"
)

func TestFileList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client, err := pikpak.NewClient(os.Getenv("PIKPAK_USERNAME"), os.Getenv("PIKPAK_PASSWORD"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = client.Login()
	assert.NoError(t, err)

	files, err := client.FileList(os.Getenv("PIKPAK_PATH"))
	assert.NoError(t, err)
	assert.NotEmpty(t, files)
	t.Log(len(files))
	t.Log(files)
}

func TestLink(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client, err := pikpak.NewClient(os.Getenv("PIKPAK_USERNAME"), os.Getenv("PIKPAK_PASSWORD"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	err = client.Login()
	assert.NoError(t, err)

	link, err := client.Link(os.Getenv("PIKPAK_FILE"))
	assert.NoError(t, err)
	assert.NotEmpty(t, link)
	t.Log(link)
}
