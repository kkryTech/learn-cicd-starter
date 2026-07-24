package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key-123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if apiKey != "test-api-key-123" {
		t.Errorf("expected 'test-api-key-123', got '%s'", apiKey)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error for missing header, got nil")
	}
	if apiKey != "" {
		t.Errorf("expected empty apiKey, got '%s'", apiKey)
	}
}

func TestGetAPIKey_InvalidFormat(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer invalid-token")

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error for invalid format, got nil")
	}
	if apiKey != "" {
		t.Errorf("expected empty apiKey, got '%s'", apiKey)
	}
}
