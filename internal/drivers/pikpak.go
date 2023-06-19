package drivers

import (
	"strconv"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/pikpak"
	"github.com/samber/lo"
)

type PikPak struct {
	client *pikpak.Client
}

func NewPikPak(username, password string) (*PikPak, error) {
	client, err := pikpak.NewClient(username, password)
	if err != nil {
		return nil, err
	}
	return &PikPak{client: client}, nil
}

func (s *PikPak) Name() string {
	return "PikPak"
}

func (s *PikPak) Init() error {
	return s.client.Login()
}

func (s *PikPak) IsSupported() IsSupported {
	return IsSupported{
		FileList:    true,
		Link:        true,
		Share:       false,
		ShareCancel: false,
		ShareList:   false,
	}
}

func (s *PikPak) FileList(id string) ([]FileDirectory, error) {
	list, err := s.client.FileList(id)
	if err != nil {
		return nil, err
	}
	files := lo.Map(list, func(f pikpak.File, i int) FileDirectory {
		fd := FileDirectory{
			ID:    f.ID,
			Name:  f.Name,
			Path:  id,
			IsDir: f.Kind == "drive#folder",
			Size:  -1,
		}
		if f.Size != "" {
			size, err := strconv.ParseInt(f.Size, 10, 64)
			if err == nil {
				fd.Size = size
			}
		}
		return fd
	})
	return files, nil
}

func (s *PikPak) Link(id string) (*DownloadInfo, error) {
	du, err := s.client.Link(id)
	if err != nil {
		return nil, err
	}
	return &DownloadInfo{Url: du.Url, Filename: du.Filename}, nil
}

func (s *PikPak) Share(_, _ string, _ time.Duration, _ utils.MS) (*ShareInfo, error) {
	return nil, ErrNotImplemented
}

func (s *PikPak) ShareCancel(_ []string) error {
	return ErrNotImplemented
}

func (s *PikPak) ShareList(_ int) ([]ShareInfo, error) {
	return nil, ErrNotImplemented
}
