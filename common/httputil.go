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
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 8)) //设置发送接受数据超时
				return conn, nil
			},
			//ResponseHeaderTimeout: time.Second * 5,
		},
	}
}
