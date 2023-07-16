package flights

import "net/http"

type HttpClientMock struct {
	Response func() (*http.Response, error)
}

func (c *HttpClientMock) Do(req *http.Request) (retres *http.Response, reterr error) {
	return c.Response()
}
