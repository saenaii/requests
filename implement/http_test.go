package implement

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestImpl_Get(t *testing.T) {
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
			r, err := h.Get(mockServer.URL, nil)
			if c.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, r.Text(), c.expect)
		})
	}
}

func TestImpl_Post(t *testing.T) {
	testTable := []struct {
		name          string
		handler       func(w http.ResponseWriter, r *http.Request)
		header        map[string]string
		contentMap    map[string]string
		contentString string
		wantErr       bool
	}{
		{
			name: "happy path",
			handler: func(w http.ResponseWriter, r *http.Request) {
				b, _ := ioutil.ReadAll(r.Body)
				w.Write(b)
			},
			header: map[string]string{
				"content-type": "application/x-www-form-urlencoded",
			},
			contentMap: map[string]string{
				"name": "abc",
				"age":  "20",
			},
			wantErr: false,
		},
	}

	for _, c := range testTable {
		t.Run(c.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(c.handler))

			h := NewClient()
			h.SetTimeout(time.Second)
			payload := BuildQuery(c.contentMap)
			r, err := h.Post(mockServer.URL, c.header, strings.NewReader(payload))
			if c.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, payload, r.Text())
		})
	}
}
