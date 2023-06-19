package dproxy

type Client struct {
	signKey  string
	proxyUrl string
}

func NewClient(signKey, proxyUrl string) *Client {
	if proxyUrl == "" {
		return nil
	}
	if proxyUrl[len(proxyUrl)-1] == '/' {
		proxyUrl = proxyUrl[:len(proxyUrl)-1]
	}
	return &Client{
		signKey:  signKey,
		proxyUrl: proxyUrl,
	}
}
