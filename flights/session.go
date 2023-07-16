package flights

import (
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Session struct {
	client Client
}

func New() *Session {
	return &Session{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
