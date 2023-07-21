package flights

import (
	"net/http"
	"sync"
	"time"
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
	Do(req *http.Request) (*http.Response, error)
}

type Session struct {
	Cities Map[string, string]

	client HttpClient
}

func New() *Session {
	return &Session{
		Cities: Map[string, string]{},
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
