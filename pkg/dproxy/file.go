package dproxy

import (
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/gizmo-ds/panshare/internal/utils"
)

func (s *Client) Link(filename string, expire time.Duration, params utils.MS) (string, error) {
	params["_filename"] = filename
	params["expire"] = strconv.FormatInt(time.Now().Add(expire).Unix(), 10)
	params["sign"] = s.sign(params)
	delete(params, "path")

	u, err := url.Parse(s.proxyUrl + filename)
	if err != nil {
		return "", err
	}
	u.RawQuery = utils.ParamsToValues(params).Encode()
	return u.String(), nil
}

func (s *Client) sign(params utils.MS) string {
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := s.signKey
	for _, k := range keys {
		str += k + params[k]
	}
	str += s.signKey
	str = s.signKey + utils.ToMD5String([]byte(str)) + s.signKey
	str = utils.ToMD5String([]byte(str))
	return str
}
