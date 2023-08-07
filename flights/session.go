// Package flights is a client library for the Google Flight API.
package flights

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/ratelimit"
)

// Map is safe for concurrent use by multiple goroutines. This is a wrapper around
// [sync.Map]. The purpose of it is to avoid type assertions when loading elements
// from the Map.
type Map[K comparable, V any] struct {
	m sync.Map
}

// Load returns the value stored in the map for a key, or zero if no value is present.
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
// It is safe for concurrent use by multiple goroutines. (Concurrent example: [github.com/krisukox/google-flights-api/examples/example3])
type Session struct {
	Cities Map[string, string] // Map which acts like a cache: city name -> abbravated city names

	client httpClient
	rl     ratelimit.Limiter
	cookie string
	logger *log.Logger
}

func customRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp == nil {
			return true, fmt.Errorf("response is nil")
		}

		if resp.StatusCode != http.StatusOK {
			return true, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}
}

func New() *Session {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	// client.Logger = log.New(os.Stdout, "", 0)
	client.Logger = nil
	client.CheckRetry = customRetryPolicy()
	client.RetryWaitMin = time.Second
	// client.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {

	// }
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	return &Session{
		Cities: Map[string, string]{},
		client: client,
		rl:     ratelimit.New(1),
		cookie: strings.Split(res.Header["Set-Cookie"][2], ";")[0],
		logger: log.New(os.Stdout, "", 0),
	}
}
