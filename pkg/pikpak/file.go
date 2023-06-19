package pikpak

import (
	"errors"

	"github.com/gizmo-ds/panshare/internal/utils"
)

func (s *Client) FileList(id string) ([]File, error) {
	r := make([]File, 0)

	pageToken := "first"
	for pageToken != "" {
		if pageToken == "first" {
			pageToken = ""
		}

		var result struct {
			ErrResult
			Files
		}
		_, err := s.client.R().
			SetAuthToken(s.token.AccessToken).
			SetQueryParams(utils.MS{
				"parent_id":      id,
				"thumbnail_size": "SIZE_LARGE",
				"with_audit":     "true",
				"limit":          "100",
				"filters":        `{"phase":{"eq":"PHASE_TYPE_COMPLETE"},"trashed":{"eq":false}}`,
				"page_token":     pageToken,
			}).
			SetResult(&result).
			Get("https://api-drive.mypikpak.com/drive/v1/files")
		if err != nil {
			return nil, err
		}

		if result.ErrorCode != 0 {
			if result.ErrorCode == 16 {
				err = s.refreshToken()
				if err != nil {
					return nil, err
				}
				return s.FileList(id)
			}
			return nil, errors.New(result.Error)
		}

		pageToken = result.NextPageToken
		r = append(r, result.Files.Files...)
	}
	return r, nil
}

func (s *Client) Link(id string) (DownloadUrl, error) {
	var u DownloadUrl

	var result struct {
		ErrResult
		File
	}
	_, err := s.client.R().
		SetAuthToken(s.token.AccessToken).
		SetQueryParams(utils.MS{"usage": "FETCH"}).
		SetResult(&result).
		Get("https://api-drive.mypikpak.com/drive/v1/files/" + id)
	if err != nil {
		return u, err
	}

	if result.ErrorCode != 0 {
		if result.ErrorCode == 16 {
			err = s.refreshToken()
			if err != nil {
				return u, err
			}
			return s.Link(id)
		}
		return u, errors.New(result.Error)
	}
	u.Filename = result.Name
	u.Url = result.WebContentLink

	if len(result.Medias) > 0 && result.Medias[0].Link.Url != "" {
		u.Url = result.Medias[0].Link.Url
	}
	return u, nil
}
