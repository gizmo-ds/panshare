package drivers

import (
	"errors"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
)

type (
	Base interface {
		Name() string
		Init() error
		IsSupported() IsSupported
		FileList(string) ([]FileDirectory, error)
		Link(string) (*DownloadInfo, error)
		Share(key string, pwd string, exp time.Duration, params utils.MS) (*ShareInfo, error)
		ShareCancel([]string) error
		ShareList(int) ([]ShareInfo, error)
	}

	IsSupported struct {
		FileList    bool
		Link        bool
		Share       bool
		ShareCancel bool
		ShareList   bool
	}
	FileDirectory struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Path  string `json:"path"`
		Size  int64  `json:"size"`
		IsDir bool   `json:"is_dir"`
	}
	ShareInfo struct {
		Url string `json:"url"`
		Pwd string `json:"pwd"`
		Exp int64  `json:"exp"`
	}
	DownloadInfo struct {
		Url      string `json:"url"`
		Filename string `json:"filename"`
	}
)

var ErrNotImplemented = errors.New("not implemented")
