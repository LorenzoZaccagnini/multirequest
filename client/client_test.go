package client

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/rs/zerolog"
)

func TestCheckRequest(t *testing.T) {
	//servers

	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	invalidServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	testCases := []struct {
		name          string
		url           string
		expectedValid bool
	}{
		{
			name:          "Valid URL",
			url:           "https://validurl.com",
			expectedValid: true,
		},
		{
			name:          "Invalid URL",
			url:           "https://notvalidurl.com",
			expectedValid: false,
		},
		{
			name:          "Valid URL with mock server",
			url:           validServer.URL,
			expectedValid: true,
		},
		{
			name:          "Invalid URL with mock server",
			url:           invalidServer.URL,
			expectedValid: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			valid := CheckRequest(tt.url, zerolog.Nop(), &wg)
			if valid != tt.expectedValid {
				t.Errorf("isValidStatus(%q) = %v, want %v", tt.url, valid, tt.expectedValid)
			}
			wg.Wait()
		})
	}
}
