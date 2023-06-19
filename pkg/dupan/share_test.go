package dupan_test

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gizmo-ds/panshare/pkg/dupan"

	"github.com/stretchr/testify/assert"
)

func TestShareList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client := dupan.NewClient()
	assert.NotNil(t, client)
	client.Setup(os.Getenv("BDUSS"), "")

	list, err := client.ShareList(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	for _, info := range list {
		t.Logf("%+v\n", info)
	}
}

func TestShareURLInfo(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	shareID, err := strconv.ParseInt(os.Getenv("TEST_SHARE_ID"), 10, 64)
	assert.NoError(t, err)

	client := dupan.NewClient()
	assert.NotNil(t, client)
	client.Setup(os.Getenv("BDUSS"), "")

	info, err := client.ShareURLInfo(shareID)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.NotEmpty(t, info.ShortURL)
	t.Log(info)
}

func TestShare(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	sharePaths := strings.Split(os.Getenv("TEST_SHARE_PATHS"), ",")
	t.Log(sharePaths)

	client := dupan.NewClient()
	assert.NotNil(t, client)
	client.Setup(os.Getenv("BDUSS"), "")

	info, err := client.Share(sharePaths, "", 7)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Link)
	t.Log(info)
}

func TestShareCancel(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	shareID, err := strconv.ParseInt(os.Getenv("TEST_SHARE_ID"), 10, 64)
	assert.NoError(t, err)

	client := dupan.NewClient()
	assert.NotNil(t, client)
	client.Setup(os.Getenv("BDUSS"), "")

	err = client.ShareCancel([]int64{shareID})
	assert.NoError(t, err)
}
