package util

import (
	"net/http"
	"time"
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 5,
		MaxConnsPerHost:     512,
		IdleConnTimeout:     time.Second * 10,
	},
}

func GetDefaultClient() *http.Client {
	return defaultClient
}
