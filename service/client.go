package service

import "sync"

var (
	once     sync.Once
	instance *Client
)

type Client struct {
}

func New() *Client {
	once.Do(func() {
		instance = &Client{}
	})
	return instance
}
