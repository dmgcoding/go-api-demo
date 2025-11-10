package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"api_demo/internal/api"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedMessage string
		expectedError  string
	}{
		{
			name:           "valid name starting with A",
			queryParam:     "name=Alice",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello Alice",
		},
		{
			name:           "valid name starting with M",
			queryParam:     "name=Mike",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello Mike",
		},
		{
			name:           "valid name starting with lowercase a",
			queryParam:     "name=alice",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello alice",
		},
		{
			name:           "valid name starting with lowercase m",
			queryParam:     "name=mike",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello mike",
		},
		{
			name:           "invalid name starting with N",
			queryParam:     "name=Nancy",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "invalid name starting with Z",
			queryParam:     "name=Zoe",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "invalid name starting with lowercase n",
			queryParam:     "name=nancy",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "invalid name starting with lowercase z",
			queryParam:     "name=zoe",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "empty name parameter",
			queryParam:     "name=",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "missing name parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "valid single character A",
			queryParam:     "name=A",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello A",
		},
		{
			name:           "invalid single character N",
			queryParam:     "name=N",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
		{
			name:           "valid name with special characters",
			queryParam:     "name=Alice123",
			expectedStatus: http.StatusOK,
			expectedMessage: "Hello Alice123",
		},
		{
			name:           "invalid name with special characters",
			queryParam:     "name=Nancy123",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid Input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/hello-world?"+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(api.HelloWorldHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var response struct {
					Message string `json:"message"`
				}
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Errorf("failed to unmarshal response: %v", err)
				}
				if response.Message != tt.expectedMessage {
					t.Errorf("handler returned wrong message: got %v want %v",
						response.Message, tt.expectedMessage)
				}
			} else {
				var response struct {
					Error string `json:"error"`
				}
				if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
					t.Errorf("failed to unmarshal response: %v", err)
				}
				if response.Error != tt.expectedError {
					t.Errorf("handler returned wrong error: got %v want %v",
						response.Error, tt.expectedError)
				}
			}

			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("handler returned wrong content type: got %v want application/json",
					contentType)
			}
		})
	}
}

