package main

import i "httpclient/iface"

type ClientConfig struct {
	BaseURL string
}

func (C *ClientConfig) GetBaseUrl() string {
	return C.BaseURL
}

func (C *ClientConfig) Clone() i.Client {
	return &ClientConfig{
		BaseURL: C.BaseURL,
	}
}
