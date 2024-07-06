package keyvalue

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			wantBody:   "Invalid Method Request\n",
		},
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/{key}", GetValueByKey)

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
