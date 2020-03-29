package requests

import "github.com/saenaii/requests/implement"

type API interface {
	GET(url string, header map[string]string) (*implement.Response, error)
}
