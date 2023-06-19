package dupan

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Mobile/15E148 Safari/604.1"
	PAppid           = "778750"
)

var (
	DefaultBaseURL = string([]byte{0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x61, 0x6e, 0x2e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x2e, 0x63, 0x6f, 0x6d})
	PUseragent     = string([]byte{0x73, 0x6f, 0x66, 0x74, 0x78, 0x6d, 0x3b, 0x6e, 0x65, 0x74, 0x64, 0x69, 0x73, 0x6b})
	PBaseurl       = string([]byte{0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70, 0x63, 0x73, 0x2e, 0x62, 0x61, 0x69, 0x64, 0x75, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x32, 0x2e, 0x30, 0x2f, 0x70, 0x63, 0x73})
)

type Client struct {
	client *resty.Client
}

func NewClient() *Client {
	restyClient := resty.New().
		SetBaseURL(DefaultBaseURL).
		SetHeader("User-Agent", DefaultUserAgent)
	return &Client{client: restyClient}
}

func (s *Client) SetUserAgent(ua string) *Client {
	s.client.SetHeader("User-Agent", ua)
	return s
}

func (s *Client) Setup(bduss, stoken string) *Client {
	s.client.Cookies = make([]*http.Cookie, 0)
	if bduss != "" {
		s.client.SetCookie(&http.Cookie{Name: "BDUSS", Value: bduss})
	}
	if stoken != "" {
		s.client.SetCookie(&http.Cookie{Name: "STOKEN", Value: stoken})
	}
	return s
}

func (s *Client) SetDebug(debug bool) *Client {
	s.client.SetDebug(debug)
	return s
}
