package client

import "time"

type Option func(client *RcClient)

func MaxTimeout(timeout int) Option {
	return func(c *RcClient) {
		c.timeout = time.Duration(timeout) * time.Second
	}
}
