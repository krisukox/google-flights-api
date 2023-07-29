package flights

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

type Map[K comparable, V any] struct {
	m sync.Map
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

func (m *Map[K, V]) Store(key K, value V) { m.m.Store(key, value) }

type HttpClient interface {
	Do(req *retryablehttp.Request) (*http.Response, error)
}

type Session struct {
	Cities Map[string, string]

	client HttpClient
}

func CustomRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp.StatusCode != http.StatusOK {
			return true, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}
}

func New() *Session {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.Logger = nil
	client.CheckRetry = CustomRetryPolicy()

	return &Session{
		Cities: Map[string, string]{},
		client: client,
	}
}
