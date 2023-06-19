package drivers_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gizmo-ds/panshare/internal/drivers"
	"github.com/gizmo-ds/panshare/pkg/s3"
	"github.com/stretchr/testify/assert"
)

func TestS3_FileList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	d, err := drivers.NewS3(drivers.S3Options{
		Options: s3.Options{
			EndPoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			CustomHost:   os.Getenv("S3_CUSTOM_HOST"),
			RemoveBucket: os.Getenv("S3_REMOVE_BUCKET") == "true",
		},
		Bucket:    os.Getenv("S3_BUCKET"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, d)

	list, err := d.FileList("/")
	assert.NoError(t, err)
	assert.NotEmpty(t, list)

	for _, f := range list {
		t.Logf("%+v", f)
	}
}

func TestS3_Link(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	d, err := drivers.NewS3(drivers.S3Options{
		Options: s3.Options{
			EndPoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			CustomHost:   os.Getenv("S3_CUSTOM_HOST"),
			RemoveBucket: os.Getenv("S3_REMOVE_BUCKET") == "true",
		},
		Bucket:    os.Getenv("S3_BUCKET"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, d)

	files := strings.Split(os.Getenv("S3_TEST_SHARE_OBJECTS"), ",")

	for _, f := range files {
		link, err := d.Link(f)
		assert.NoError(t, err)
		assert.NotEmpty(t, link)
		t.Log(link)
	}
}
