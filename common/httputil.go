package common

import (
	"net"
	"net/http"
	"time"
)

var HttpClient *http.Client

func init() {
	HttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			//ResponseHeaderTimeout: time.Second * 5,
		},
	}
}
