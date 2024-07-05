package keyvalue

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetValueHandler(t *testing.T) {
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
			method:     http.MethodGet,
			key:        "key",
			body:       nil,
			wantStatus: http.StatusMethodNotAllowed,
			wantBody:   "Invalid Method Request\n",
		},
		{
			name:       "Empty Value",
			method:     http.MethodPost,
			key:        "key",
			body:       bytes.NewReader([]byte("")),
			wantStatus: http.StatusBadRequest,
			wantBody:   "Bad Request: missing value\n",
		},
		{
			name:       "Valid Request",
			method:     http.MethodPost,
			key:        "key",
			body:       bytes.NewReader([]byte("value")),
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"key key and keyvalue value has been added."}` + "\n",
		},
	}

	mux := &http.ServeMux{}
	mux.HandleFunc("/{key}", SetValueHandler)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/"+tt.key, tt.body)
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
