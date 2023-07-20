package flights

import (
	"net/http"
	"testing"
)

type HttpClientMock struct {
	Responses []func() (*http.Response, error)
	T         *testing.T
}

func (c *HttpClientMock) Do(req *http.Request) (retres *http.Response, reterr error) {
	if len(c.Responses) == 0 {
		c.T.Fatalf("HttpClientMock: lack of responses")
	}

	var r func() (*http.Response, error)
	r, c.Responses = c.Responses[0], c.Responses[1:]
	return r()
}
