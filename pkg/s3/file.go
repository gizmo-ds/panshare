package s3

import (
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/samber/lo"
)

func (s *Client) FileList(bucket, path string) ([]Object, error) {
	key := strings.TrimPrefix(path, "/")
	objects := make([]Object, 0)
	err := s.client.ListObjectsPages(&s3.ListObjectsInput{
		Bucket:  &bucket,
		Prefix:  &key,
		MaxKeys: aws.Int64(1000),
	}, func(page *s3.ListObjectsOutput, lp bool) bool {
		objs := lo.Map(page.Contents, func(item *s3.Object, i int) Object {
			return Object{Key: *item.Key, Size: *item.Size}
		})
		objects = append(objects, objs...)
		if len(objs) < 1000 {
			return false
		}
		return !lp
	})
	return objects, err
}

func (s *Client) FileListV2(bucket, path string) ([]Object, error) {
	key := strings.TrimPrefix(path, "/")
	objects := make([]Object, 0)
	err := s.client.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket:  &bucket,
		Prefix:  &key,
		MaxKeys: aws.Int64(1000),
	}, func(page *s3.ListObjectsV2Output, lp bool) bool {
		objs := lo.Map(page.Contents, func(item *s3.Object, i int) Object {
			return Object{Key: *item.Key, Size: *item.Size}
		})
		objects = append(objects, objs...)
		if len(objs) < 1000 {
			return false
		}
		return !lp
	})
	return objects, err
}

func (s *Client) Link(bucket, key string, expire time.Duration) (string, error) {
	key = strings.TrimPrefix(key, "/")
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	link, err := req.Presign(expire)
	if err != nil {
		return "", err
	}
	u, _ := url.Parse(link)
	if s.removeBucket {
		u.Path = strings.Replace(u.Path, "/"+bucket, "", 1)
	}
	if s.customHost != "" {
		u.Host = s.customHost
	}
	return u.String(), nil
}
