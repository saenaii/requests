package implement

import (
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout = time.Second * 5
)

type Impl struct {
	client *http.Client
}

func NewClient() *Impl {
	return &Impl{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// GET request
func (h *Impl) GET(url string, header map[string]string) (*Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &Response{}, err
	}
	addHeader(request, header)

	r, err := h.client.Do(request)
	if err != nil {
		return &Response{}, err
	}
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &Response{}, err
	}

	return &Response{
		Content:    content,
		StatusCode: r.StatusCode,
	}, nil
}

// SetTimeout ...
func (h *Impl) SetTimeout(duration time.Duration) {
	h.client.Timeout = duration
}

func addHeader(r *http.Request, header map[string]string) {
	if header == nil {
		return
	}
	for k, v := range header {
		r.Header.Add(k, v)
	}
}
