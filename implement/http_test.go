package implement

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestImpl_GET(t *testing.T) {
	testTable := []struct {
		name    string
		handler func(w http.ResponseWriter, r *http.Request)
		header  map[string]string
		expect  string
		wantErr bool
	}{
		{
			name: "happy path",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello world"))
			},
			expect:  "Hello world",
			wantErr: false,
		}, {
			name: "happy path with header",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello world"))
			},
			header: map[string]string{
				"content-type": "application/x-www-form-urlencoded",
			},
			expect:  "Hello world",
			wantErr: false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(c.handler))

			h := NewClient()
			h.SetTimeout(time.Second)
			r, err := h.GET(mockServer.URL, nil)
			if c.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, r.Text(), c.expect)
		})
	}
}
