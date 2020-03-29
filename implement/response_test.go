package implement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_Text(t *testing.T) {
	testTable := []struct {
		name    string
		content []byte
		expect  string
	}{
		{
			name:    "happy path",
			content: []byte("hello world"),
			expect:  "hello world",
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			h := &Response{
				Content: c.content,
			}
			assert.Equal(t, c.expect, h.Text())
		})
	}
}

func TestResponse_Json(t *testing.T) {
	type response struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	testTable := []struct {
		name    string
		content []byte
		expect  response
		wantErr bool
	}{
		{
			name:    "happy path",
			content: []byte(`{"name":"abc","age":20}`),
			expect:  response{Name: "abc", Age: 20},
			wantErr: false,
		}, {
			name:    "unmarshal fail",
			content: []byte("Hello world"),
			wantErr: true,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			h := &Response{Content: c.content}
			var v response

			err := h.Json(&v)
			if c.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, c.expect, v)
		})
	}
}
