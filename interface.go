package requests

import (
	"io"

	"github.com/saenaii/requests/implement"
)

type API interface {
	Get(url string, header map[string]string) (*implement.Response, error)
	Post(url string, header map[string]string, payload io.Reader) (*implement.Response, error)
}

func NewInstance() *implement.Impl {
	return implement.NewClient()
}
