package main

import (
	"net"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

type UpstreamDialer struct {
	dialer proxy.Dialer
}

func NewUpstreamDialer(socksAddr string, timeout time.Duration) (*UpstreamDialer, error) {
	var dialer proxy.Dialer

	if socksAddr != "" {
		parsedURL, err := url.Parse(socksAddr)
		user := parsedURL.User.Username()
		password, _ := parsedURL.User.Password()
		socksDialer, err := proxy.SOCKS5(
			"tcp", parsedURL.Host,
			&proxy.Auth{User: user, Password: password},
			proxy.Direct,
		)
		if err != nil {
			return nil, err
		}
		dialer = socksDialer
	} else {
		dialer = &net.Dialer{Timeout: timeout}
	}

	return &UpstreamDialer{dialer: dialer}, nil
}

func (u *UpstreamDialer) Dial(network, addr string) (net.Conn, error) {
	return u.dialer.Dial(network, addr)
}
