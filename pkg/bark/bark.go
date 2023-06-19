package bark

import (
	"errors"
	"net/http"

	"github.com/gizmo-ds/panshare/internal/utils"
	"github.com/go-resty/resty/v2"
)

var client = resty.New().
	SetBaseURL("https://api.day.app")

func Push(key string, body utils.M) error {
	var result struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}
	body["device_key"] = key
	_, err := client.SetDebug(true).R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&result).
		Post("push")
	if err != nil {
		return err
	}
	if result.Code != http.StatusOK {
		return errors.New(result.Message)
	}
	return nil
}
