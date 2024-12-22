package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"secretary/alpha/api"
)

func mockMiddleware(w http.ResponseWriter, r *http.Request, secure ...bool) bool {
	// Mock Middleware that simulates authentication
	w.Header().Set("Content-Type", "application/json")
	return true // Always allow for testing
}

func TestAskAPI(t *testing.T) {
	originalMiddleware := api.MiddlewareWrapper
	api.MiddlewareWrapper = mockMiddleware
	defer func() {
		api.MiddlewareWrapper = originalMiddleware
	}()

	tests := []struct {
		name           string
		method         string
		body           map[string]interface{}
		queryParams    string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Create Ask - Valid Input",
			method:         "POST",
			body:           map[string]interface{}{"what": "Task1", "reason": "For Testing", "reviewer": "user1"},
			expectedStatus: 201,
			expectedBody:   map[string]interface{}{"ask_data": "asksfor created successfully"},
		},
		{
			name:           "Create Ask - Missing Field",
			method:         "POST",
			body:           map[string]interface{}{"what": "Task1", "reviewer": "user1"},
			expectedStatus: 400,
			expectedBody:   map[string]interface{}{"error": "invalid data"},
		},
		{
			name:           "Get All Asks - No UUID",
			method:         "GET",
			queryParams:    "",
			expectedStatus: 200,
			expectedBody:   map[string]interface{}{"ask_data": []interface{}{}},
		},
		{
			name:           "Get Ask by UUID - Valid UUID",
			method:         "GET",
			queryParams:    "?uuid=12345",
			expectedStatus: 200,
			expectedBody:   map[string]interface{}{"ask_data": map[string]interface{}{"uuid": "12345"}},
		},
		{
			name:           "Invalid Method",
			method:         "PUT",
			expectedStatus: 405,
			expectedBody:   map[string]interface{}{"error": "method not allowed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != nil {
				bodyBytes, _ := json.Marshal(tt.body)
				req = httptest.NewRequest(tt.method, "/api/asksfor"+tt.queryParams, bytes.NewReader(bodyBytes))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, "/api/asksfor"+tt.queryParams, nil)
			}

			recorder := httptest.NewRecorder()

			// Simulate the API call using the mock middleware
			api.AskAPI(recorder, req)

			// Check for expected status code
			if recorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}

			// Unmarshal and check response body
			var respBody map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &respBody)
			for key, expectedValue := range tt.expectedBody {
				if respBody[key] != expectedValue {
					t.Errorf("expected body key %s to be %v, got %v", key, expectedValue, respBody[key])
				}
			}
		})
	}
}

func TestAskNoAuthAPI(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           map[string]interface{}
		queryParams    string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Create Ask - Valid Input",
			method:         "POST",
			body:           map[string]interface{}{"what": "Task1", "reason": "For Testing", "reviewer": "user1"},
			expectedStatus: 401, //201
			expectedBody:   nil,
		},
		{
			name:           "Create Ask - Missing Field",
			method:         "POST",
			body:           map[string]interface{}{"what": "Task1", "reviewer": "user1"},
			expectedStatus: 401, //400
			expectedBody:   nil,
		},
		{
			name:           "Get All Asks - No UUID",
			method:         "GET",
			queryParams:    "",
			expectedStatus: 401, //200
			expectedBody:   nil,
		},
		{
			name:           "Get Ask by UUID - Valid UUID",
			method:         "GET",
			queryParams:    "?uuid=12345",
			expectedStatus: 401, //200
			expectedBody:   nil,
		},
		{
			name:           "Invalid Method",
			method:         "PUT",
			expectedStatus: 401, //405
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != nil {
				bodyBytes, _ := json.Marshal(tt.body)
				req = httptest.NewRequest(tt.method, "/api/asksfor"+tt.queryParams, bytes.NewReader(bodyBytes))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, "/api/asksfor"+tt.queryParams, nil)
			}

			recorder := httptest.NewRecorder()

			api.AskAPI(recorder, req)

			if recorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}

			var respBody map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &respBody)
			for key, expectedValue := range tt.expectedBody {
				if respBody[key] != expectedValue {
					t.Errorf("expected body key %s to be %v, got %v", key, expectedValue, respBody[key])
				}
			}
		})
	}
}
