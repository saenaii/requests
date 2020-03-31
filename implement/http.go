package implement

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
func (h *Impl) Get(url string, header, query map[string]string) (*Response, error) {
	if query != nil {
		url += "?" + BuildQuery(query)
	}
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

func (h *Impl) Post(url string, header map[string]string, payload io.Reader) (*Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return &Response{}, nil
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

func BuildQuery(content map[string]string) string {
	var queryString string
	for k, v := range content {
		queryString += k + "=" + v + "&"
	}
	return strings.TrimRight(queryString, "&")
}

func BuildFormData(content map[string]string) string {
	return ""
}
