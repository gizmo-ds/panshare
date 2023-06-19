package dproxy_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/dproxy"
	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	files := strings.Split(os.Getenv("DPROXY_SHARE_FILES"), ",")

	client := dproxy.NewClient(os.Getenv("DPROXY_SIGN_KEY"), os.Getenv("DPROXY_PROXY_URL"))
	assert.NotNil(t, client)

	for _, name := range files {
		link, err := client.Link(name, time.Hour*24*7, utils.MS{
			"bucket": "ams1",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, link)
		t.Log(link)
	}
}
