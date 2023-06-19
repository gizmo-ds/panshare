package drivers

import (
	"path/filepath"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/dproxy"
)

type DProxy struct {
	client *dproxy.Client
	params utils.MS
}

func NewDProxy() *DProxy {
	return &DProxy{}
}

func (s *DProxy) Name() string {
	return "DProxy"
}

func (s *DProxy) Init() error {
	return nil
}

func (s *DProxy) IsSupported() IsSupported {
	return IsSupported{
		FileList:    false,
		Link:        true,
		Share:       false,
		ShareCancel: false,
		ShareList:   false,
	}
}

func (s *DProxy) FileList(_ string) ([]FileDirectory, error) {
	return nil, ErrNotImplemented
}

func (s *DProxy) Link(key string) (*DownloadInfo, error) {
	u, err := s.client.Link(key, time.Hour*3, s.params)
	if err != nil {
		return nil, err
	}
	return &DownloadInfo{Url: u, Filename: filepath.Base(key)}, nil
}

func (s *DProxy) Share(_, _ string, _ time.Duration, _ utils.MS) (*ShareInfo, error) {
	return nil, ErrNotImplemented
}

func (s *DProxy) ShareCancel(_ []string) error {
	return ErrNotImplemented
}

func (s *DProxy) ShareList(_ int) ([]ShareInfo, error) {
	return nil, ErrNotImplemented
}
