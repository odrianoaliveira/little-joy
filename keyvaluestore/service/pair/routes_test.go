package pair

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSetValueHandler(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		contentType string
		key         string
		body        io.Reader
		wantStatus  int
		wantBody    string
	}{
		{
			name:        "Invalid Method",
			method:      http.MethodPost,
			key:         "key",
			body:        nil,
			contentType: "text/plain; charset=utf-8",
			wantStatus:  http.StatusMethodNotAllowed,
			wantBody:    "Method not allowed\n",
		},
		{
			name:        "Empty Value",
			method:      http.MethodPut,
			key:         "key",
			contentType: "text/plain; charset=utf-8",
			body:        bytes.NewReader([]byte("")),
			wantStatus:  http.StatusBadRequest,
			wantBody:    "Bad Request: missing value\n",
		},
		{
			name:        "Valid Request",
			method:      http.MethodPut,
			key:         "a key",
			contentType: "application/json",
			body:        bytes.NewReader([]byte("a value")),
			wantStatus:  http.StatusOK,
			wantBody:    `{"key":"a key","value":"a value"}` + "\n",
		},
	}
	mux := &http.ServeMux{}
	handler := NewHandler()
	handler.RegisterRoutes(mux)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keyEsc := url.PathEscape(tt.key)
			req := httptest.NewRequest(tt.method, "/pair/"+keyEsc, tt.body)
			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			res := rr.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.wantStatus, res.StatusCode)
			assert.Equal(t, tt.wantBody, string(body))
			assert.Equal(t, tt.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func TestGetValueHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		key        string
		body       io.Reader
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Invalid Method",
			method:     http.MethodPost,
			key:        "key",
			body:       nil,
			wantStatus: http.StatusMethodNotAllowed,
			wantBody:   "Method not allowed\n",
		},
		{
			name:       "Valid Request",
			method:     http.MethodGet,
			key:        "key",
			body:       nil,
			wantStatus: http.StatusOK,
			wantBody:   `{"key":"key","value":"a value"}` + "\n",
		},
	}
	mux := &http.ServeMux{}
	handler := NewHandler()
	handler.RegisterRoutes(mux)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/pair/"+tt.key, tt.body)
			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			res := rr.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.wantStatus, res.StatusCode)
			assert.Equal(t, tt.wantBody, string(body))
		})
	}
}
