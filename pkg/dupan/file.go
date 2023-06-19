package dupan

import (
	"errors"

	"github.com/gizmo-ds/panshare/internal/utils"
)

func (s *Client) FileList(path, order, by string) (FileDirectoryList, error) {
	var result struct {
		List FileDirectoryList `json:"list"`
		PResultErr
	}
	_, err := s.client.R().
		SetQueryParams(utils.MS{
			"method": "list",
			"path":   path,
			"by":     by,
			"order":  order,
			"limit":  "0-2147483647",
			"app_id": PAppid,
		}).
		SetHeader("User-Agent", PUseragent).
		SetResult(&result).
		ForceContentType("application/json; charset=utf-8").
		Get(PBaseurl + "/file")
	if err != nil {
		return nil, err
	}
	if result.ErrorCode != 0 {
		return nil, errors.New(result.ErrorMsg)
	}
	return result.List, nil
}
