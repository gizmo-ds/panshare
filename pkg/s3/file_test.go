package s3_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gizmo-ds/panshare/pkg/s3"

	"github.com/stretchr/testify/assert"
)

func TestFileList(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client, err := s3.NewClient(
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		s3.Options{
			EndPoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			CustomHost:   os.Getenv("S3_CUSTOM_HOST"),
			RemoveBucket: os.Getenv("S3_REMOVE_BUCKET") == "true",
		})
	assert.NoError(t, err)

	bucket := os.Getenv("S3_BUCKET")

	list, err := client.FileList(bucket, "/")
	assert.NoError(t, err)
	assert.NotNil(t, list)
	t.Log(len(list))
	t.Log(list)
}

func TestFileListV2(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client, err := s3.NewClient(
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		s3.Options{
			EndPoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			CustomHost:   os.Getenv("S3_CUSTOM_HOST"),
			RemoveBucket: os.Getenv("S3_REMOVE_BUCKET") == "true",
		})
	assert.NoError(t, err)

	bucket := os.Getenv("S3_BUCKET")

	list, err := client.FileListV2(bucket, "/")
	assert.NoError(t, err)
	assert.NotNil(t, list)
	t.Log(len(list))
	t.Log(list)
}

func TestLink(t *testing.T) {
	err := LoadEnv()
	assert.NoError(t, err)

	client, err := s3.NewClient(
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		s3.Options{
			EndPoint:     os.Getenv("S3_ENDPOINT"),
			Region:       os.Getenv("S3_REGION"),
			CustomHost:   os.Getenv("S3_CUSTOM_HOST"),
			RemoveBucket: os.Getenv("S3_REMOVE_BUCKET") == "true",
		})
	assert.NoError(t, err)

	bucket := os.Getenv("S3_BUCKET")
	files := strings.Split(os.Getenv("S3_TEST_SHARE_OBJECTS"), ",")

	for _, f := range files {
		link, err := client.Link(bucket, f, time.Hour*24*7)
		assert.NoError(t, err)
		assert.NotEmpty(t, link)
		t.Log(link)
	}
}
