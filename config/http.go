package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewHTTPClient returns a http.Client with sane timeouts configured
func NewHTTPClient() *http.Client {
	var tr = &http.Transport{
		// Disable Certificate Checking
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: 30 * time.Second,
		// Connection timeout = 5s
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		// TLS Handshake Timeout = 5s
		TLSHandshakeTimeout: 10 * time.Second,
	}
	// HTTP Timeout = 10s
	return &http.Client{Timeout: 30 * time.Second, Transport: tr}
}
