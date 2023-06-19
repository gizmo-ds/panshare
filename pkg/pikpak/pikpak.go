package pikpak

import (
	"errors"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/go-resty/resty/v2"
)

const (
	clientID     = "YNxT9w7GMdWvEOKa"
	clientSecret = "dbw2OtmVEeuUvIptb1Coyg"
)

type Client struct {
	client *resty.Client

	token TokenResult

	username string
	password string
}

func NewClient(username, password string) (*Client, error) {
	c := &Client{
		client:   resty.New(),
		username: username,
		password: password,
	}
	return c, nil
}

func (s *Client) Login() error {
	var result struct {
		ErrResult
		TokenResult
	}
	_, err := s.client.R().
		SetBody(utils.MS{
			"captcha_token": "",
			"client_id":     clientID,
			"client_secret": clientSecret,
			"username":      s.username,
			"password":      s.password,
		}).
		SetResult(&result).
		Post("https://user.mypikpak.com/v1/auth/signin")
	if err != nil {
		return err
	}
	if result.ErrorCode != 0 {
		return errors.New(result.Error)
	}
	s.token = result.TokenResult
	return nil
}

func (s *Client) refreshToken() error {
	var result struct {
		ErrResult
		TokenResult
	}
	_, err := s.client.SetDebug(true).R().
		SetBody(utils.MS{
			"client_id":     clientID,
			"client_secret": clientSecret,
			"grant_type":    "refresh_token",
			"refresh_token": s.token.RefreshToken,
		}).
		SetResult(&result).
		Post("https://user.mypikpak.com/v1/auth/signin")
	if err != nil {
		return err
	}
	if result.ErrorCode != 0 {
		if result.ErrorCode == 4126 {
			return s.Login()
		}
		return errors.New(result.Error)
	}
	s.token = result.TokenResult
	return nil
}
