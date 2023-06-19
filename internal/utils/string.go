package utils

import "net/url"

func ParamsToValues(p MS) url.Values {
	values := url.Values{}
	for k, v := range p {
		values.Add(k, v)
	}
	return values
}
