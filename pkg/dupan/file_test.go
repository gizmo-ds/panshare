package dupan_test

import (
	"os"
	"testing"

	"github.com/gizmo-ds/panshare/pkg/dupan"

	"github.com/stretchr/testify/assert"
)

func TestFileList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client := dupan.NewClient()
	assert.NotNil(t, client)
	client.Setup(os.Getenv("BDUSS"), "")

	list, err := client.FileList("/", "desc", "time")
	assert.NoError(t, err)
	assert.NotNil(t, list)
	for _, d := range list {
		t.Logf("%+v", d)
	}
}
