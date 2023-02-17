package ghc

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const DefaultUserAgent string = "Golang Http Client v1.0"
const DefaultRequestMethod string = "GET"
const DefaultRequestTimeout int64 = 10

type GhcClient struct {
	Url          string
	Method       string
	Client       *http.Client
	Header       map[string]string
	Cookie       map[string]string
	Timeout      int64
	Keepalive    bool
	LastResponse *GhcResponse
}

type GhcResponse struct {
	Header map[string][]string
	Body   []byte
	Status int
}

func NewClient() *GhcClient {
	return &GhcClient{
		Url:          "",
		Method:       DefaultRequestMethod,
		Header:       make(map[string]string),
		Cookie:       make(map[string]string),
		Timeout:      DefaultRequestTimeout,
		Keepalive:    true,
		LastResponse: &GhcResponse{},
	}
}

func (gc *GhcClient) SetUrl(u string) *GhcClient {
	gc.Url = u
	return gc
}

func (gc *GhcClient) SetMethod(m string) *GhcClient {
	gc.Method = m
	return gc
}

func (gc *GhcClient) SetTimeout(t int64) *GhcClient {
	gc.Timeout = t
	return gc
}

func (gc *GhcClient) SetKeepalive(k bool) *GhcClient {
	gc.Keepalive = k
	return gc
}

func (gc *GhcClient) Request(body io.Reader) (*http.Request, error) {
	return http.NewRequest(gc.Method, gc.Url, body)
}

func (gc *GhcClient) Do() (*http.Response, error) {
	gc.Client = &http.Client{
		Transport: nil,
		Jar:       nil,
		Timeout:   5 * time.Second,
	}
	req, err := gc.Request(nil)
	if err != nil {
		return nil, err
	}
	return gc.Client.Do(req)
}

func (gc *GhcClient) DoAll() ([]byte, error) {
	do, err := gc.Do()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}
	gc.LastResponse.Body = body
	gc.LastResponse.Header = do.Header
	gc.LastResponse.Status = do.StatusCode
	return body, nil
}

func (gc *GhcClient) GetResponseHeader() map[string][]string {
	return gc.LastResponse.Header
}

func (gc *GhcClient) GetResponseBody() []byte {
	return gc.LastResponse.Body
}

func (gc *GhcClient) GetResponseBodyString() string {
	return string(gc.LastResponse.Body)
}

func (gc *GhcClient) GetResponseStatus() int {
	return gc.LastResponse.Status
}
