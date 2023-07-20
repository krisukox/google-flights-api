package flights

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type HttpClientMock struct {
	Responses []func() (*http.Response, error)
	T         *testing.T
}

func newHttpClientMock(t *testing.T, respPaths ...string) (*HttpClientMock, error) {
	responses := []func() (*http.Response, error){}
	for _, respPath := range respPaths {
		respFile, err := os.Open(respPath)
		if err != nil {
			return nil, err
		}

		byteValue, err := ioutil.ReadAll(respFile)
		if err != nil {
			return nil, err
		}

		body := io.NopCloser(bytes.NewReader(byteValue))

		responses = append(responses, func() (*http.Response, error) {
			return &http.Response{
				Body: body,
			}, nil
		})
	}

	return &HttpClientMock{responses, t}, nil
}

func (c *HttpClientMock) Do(req *http.Request) (retres *http.Response, reterr error) {
	if len(c.Responses) == 0 {
		c.T.Fatalf("HttpClientMock: lack of responses")
	}

	var r func() (*http.Response, error)
	r, c.Responses = c.Responses[0], c.Responses[1:]
	return r()
}
