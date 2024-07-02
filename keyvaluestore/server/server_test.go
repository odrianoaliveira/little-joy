package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePairHandler(t *testing.T) {

	tests := []struct {
		name              string
		method            string
		body              interface{}
		expectedHttpStaus int
		expectedBody      string
	}{
		{
			name:              "Valid Request",
			method:            http.MethodPost,
			body:              KeyValuePayload{Key: "key", Value: "value"},
			expectedHttpStaus: http.StatusOK,
			expectedBody:      `{"message":"key key and value value has been added."}`,
		},
		{
			name:              "Invalid Method",
			method:            http.MethodDelete,
			body:              nil,
			expectedHttpStaus: http.StatusMethodNotAllowed,
			expectedBody:      "Invalid Method Request",
		},
		{
			name:              "Invalid Body",
			method:            http.MethodPost,
			body:              "invalid body",
			expectedHttpStaus: http.StatusBadRequest,
			expectedBody:      "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rBody, err := json.Marshal(tt.body)
			assert.NoError(t, err)

			req, err := http.NewRequest(tt.method, "/create", bytes.NewBuffer(rBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CreatePairHandler)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedHttpStaus, rr.Code)
			assert.Equal(t, tt.expectedBody, strings.TrimSpace(rr.Body.String()))
		})
	}

}
