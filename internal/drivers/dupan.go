package drivers

import (
	"strconv"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/gizmo-ds/panshare/pkg/dupan"
	"github.com/samber/lo"
)

type DuPan struct {
	client *dupan.Client

	Order   string
	OrderBy string

	bduss  string
	stoken string
}

func NewDuPan(bduss, stoken string) *DuPan {
	return &DuPan{
		bduss:  bduss,
		stoken: stoken,
	}
}

func (s *DuPan) Name() string {
	return "DuPan"
}

func (s *DuPan) Init() error {
	s.client.Setup(s.bduss, s.stoken)
	return nil
}

func (s *DuPan) IsSupported() IsSupported {
	return IsSupported{
		FileList:    true,
		Link:        false,
		Share:       true,
		ShareCancel: true,
		ShareList:   true,
	}
}

func (s *DuPan) FileList(path string) ([]FileDirectory, error) {
	list, err := s.client.FileList(path, s.Order, s.OrderBy)
	if err != nil {
		return nil, err
	}
	files := lo.Map(list, func(f *dupan.FileDirectory, i int) FileDirectory {
		return FileDirectory{
			ID:    strconv.FormatInt(f.FsID, 10),
			Name:  f.Filename,
			Path:  f.Path,
			Size:  f.Size,
			IsDir: f.IsDirInt == 1,
		}
	})
	return files, nil
}

func (s *DuPan) Link(_ string) (*DownloadInfo, error) {
	return nil, ErrNotImplemented
}

func (s *DuPan) Share(path string, pwd string, exp time.Duration, _ utils.MS) (*ShareInfo, error) {
	info, err := s.client.Share([]string{path}, pwd, int(exp/time.Hour/24))
	if err != nil {
		return nil, err
	}
	return &ShareInfo{
		Url: info.Link,
		Exp: time.Now().Add(exp).Unix(),
		Pwd: info.Pwd,
	}, nil
}

func (s *DuPan) ShareCancel(ids []string) error {
	return s.client.ShareCancel(
		lo.Map(ids, func(v string, _ int) int64 {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return 0
			}
			return i
		}),
	)
}

func (s *DuPan) ShareList(page int) ([]ShareInfo, error) {
	list, err := s.client.ShareList(page)
	if err != nil {
		return nil, err
	}
	shares := lo.Map(list, func(info dupan.ShareRecordInfo, i int) ShareInfo {
		return ShareInfo{
			Url: info.ShortLink,
			Exp: time.Now().
				Add(time.Duration(info.ExpireTime) * time.Second).
				Unix(),
		}
	})
	return shares, nil
}
