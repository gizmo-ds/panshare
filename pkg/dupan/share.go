package dupan

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/samber/lo"
)

func (s *Client) ShareList(page int) ([]ShareRecordInfo, error) {
	var result struct {
		ErrResult
		Count int               `json:"count"`
		List  []ShareRecordInfo `json:"list,omitempty"`
	}
	_, err := s.client.R().
		SetHeaders(utils.MS{}).
		SetResult(&result).
		SetQueryParams(utils.MS{
			"page":  strconv.Itoa(page),
			"desc":  "1",
			"order": "time",
		}).
		Get("share/record")
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, errors.New(result.ShowMsg)
	}
	return result.List, nil
}

func (s *Client) ShareURLInfo(shareID int64) (*ShareURLInfo, error) {
	var result struct {
		ShareURLInfo
		ErrResult
	}
	sign := append([]byte(strconv.FormatInt(shareID, 10)), 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x75, 0x72, 0x6c, 0x69, 0x6e, 0x66, 0x6f, 0x21, 0x40, 0x23)
	_, err := s.client.R().
		SetQueryParams(utils.MS{
			"shareid": strconv.FormatInt(shareID, 10),
			"sign":    utils.ToMD5String(sign),
		}).
		SetResult(&result).
		Get("share/" + string([]byte{0x73, 0x75, 0x72, 0x6c, 0x69, 0x6e, 0x66, 0x6f, 0x69, 0x6e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64}))
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, errors.New(result.ShowMsg)
	}
	return &result.ShareURLInfo, nil
}

func (s *Client) Share(paths []string, pwd string, period int) (*Shared, error) {
	var result struct {
		Shared
		ErrResult
	}
	if len(pwd) != 4 {
		pwd = lo.RandomString(4, []rune("abcdefghijklmnopqrstuvwxyz0123456789"))
	}
	pathList, err := json.Marshal(paths)
	if err != nil {
		return nil, err
	}
	_, err = s.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(utils.MS{
			"path_list":    string(pathList),
			"schannel":     "4",
			"channel_list": "[]",
			"period":       strconv.Itoa(period),
			"pwd":          pwd,
			"share_type":   "9",
		}).
		SetResult(&result).
		Post("share/pset")
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, errors.New(result.ShowMsg)
	}
	result.Shared.Pwd = pwd
	return &result.Shared, nil
}

func (s *Client) ShareCancel(shareIds []int64) error {
	var result ErrResult
	idList, err := json.Marshal(
		lo.Map(shareIds, func(id int64, i int) string {
			return strconv.FormatInt(id, 10)
		}),
	)
	if err != nil {
		return err
	}
	_, err = s.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(utils.MS{
			"shareid_list": string(idList),
		}).
		SetResult(&result).
		Post("share/cancel")
	if err != nil {
		return err
	}
	if result.ErrNo != 0 {
		return errors.New(result.ShowMsg)
	}
	return nil
}
