package client

import (
	"github.com/aws/aws-xray-sdk-go/xray"
	"net/http"
	"time"
)

const (
	_defaultTimeout = 30
)

type RcClient struct {
	timeout time.Duration
	Http    *http.Client
}

func New(opts ...Option) *RcClient {
	client := &RcClient{timeout: _defaultTimeout}

	for _, opt := range opts {
		opt(client)
	}

	client.Http = xray.Client(&http.Client{Timeout: client.timeout * time.Second})
	return client
}
