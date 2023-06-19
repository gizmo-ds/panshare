package drivers

import (
	"path/filepath"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/s3"
	"github.com/samber/lo"
)

type S3 struct {
	client *s3.Client

	options S3Options
}
type S3Options struct {
	s3.Options
	Bucket    string
	AccessKey string
	SecretKey string
}

func NewS3(options S3Options) (*S3, error) {
	client, err := s3.NewClient(options.AccessKey, options.SecretKey, options.Options)
	if err != nil {
		return nil, err
	}
	return &S3{
		options: options,
		client:  client,
	}, nil
}

func (s *S3) Name() string {
	return "S3"
}

func (s *S3) Init() error {
	return nil
}

func (s *S3) IsSupported() IsSupported {
	return IsSupported{
		FileList:    true,
		Link:        true,
		Share:       false,
		ShareCancel: false,
		ShareList:   false,
	}
}

func (s *S3) FileList(key string) ([]FileDirectory, error) {
	list, err := s.client.FileList(s.options.Bucket, key)
	if err != nil {
		return nil, err
	}
	files := lo.Map(list, func(k s3.Object, i int) FileDirectory {
		return FileDirectory{
			ID:    k.Key,
			Name:  filepath.Base(k.Key),
			Path:  filepath.Dir(k.Key),
			Size:  k.Size,
			IsDir: false,
		}
	})
	return files, nil
}

func (s *S3) Link(key string) (*DownloadInfo, error) {
	u, err := s.client.Link(s.options.Bucket, key, time.Hour*3)
	if err != nil {
		return nil, err
	}
	return &DownloadInfo{Url: u, Filename: filepath.Base(key)}, nil
}

func (s *S3) Share(_, _ string, _ time.Duration, _ utils.MS) (*ShareInfo, error) {
	return nil, ErrNotImplemented
}

func (s *S3) ShareCancel(_ []string) error {
	return ErrNotImplemented
}

func (s *S3) ShareList(_ int) ([]ShareInfo, error) {
	return nil, ErrNotImplemented
}
