// Package flights is a client library for Google Flight API.
//
// The client uses github.com/hashicorp/go-retryablehttp under the hood.
// Every request to the Google Flight API is retried 5 times in case of error.
package flights

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

// Map is safe for concurrent use by multiple goroutines. This is a wrapper around
// [sync.Map]. The purpose of it is to avoid ambiguous type assertion when
// loading element from the Map.
type Map[K comparable, V any] struct {
	m sync.Map
}

// Load returns the value stored in the map for a key, or zero value if no value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) { m.m.Store(key, value) }

type httpClient interface {
	Do(req *retryablehttp.Request) (*http.Response, error)
}

// Session is the main type that contains all the most important functions to operate the Google Flight API.
// It is safe to use it by mutliple goroutines.
type Session struct {
	Cities Map[string, string] // Map which acts like a cache: city name -> abbravated city names

	client httpClient
}

func customRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
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
	client.CheckRetry = customRetryPolicy()

	return &Session{
		Cities: Map[string, string]{},
		client: client,
	}
}
