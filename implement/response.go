package implement

import "encoding/json"

type Response struct {
	Content    []byte
	StatusCode int
}

// Text return a string
func (h *Response) Text() string {
	return string(h.Content)
}

// Json ...
func (h *Response) Json(v interface{}) error {
	if err := json.Unmarshal(h.Content, v); err != nil {
		return err
	}
	return nil
}
