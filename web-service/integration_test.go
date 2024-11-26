package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestFibonacciEndpoint(t *testing.T) {

	router := SetupRouter()

	router.POST("/", fibonacciHandler)

	tests := []struct {
		name           string
		input          map[string]interface{}
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "valid input",
			input:          map[string]interface{}{"number": 5},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"result": []interface{}{0.0, 1.0, 1.0, 2.0, 3.0}},
		},
		{
			name:           "negative number",
			input:          map[string]interface{}{"number": -1},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Number must be non-negative"},
		},
		{
			name:           "zero",
			input:          map[string]interface{}{"number": 0},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"result": []interface{}{}},
		},
		{
			name:           "no input",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Number field is required"},
			input:          map[string]interface{}{},
		},
		{
			name:           "number key with incorrect type value",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "Invalid input"},
			input:          map[string]interface{}{"number": []int{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var responseBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}
