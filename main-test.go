package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEncryptHandler(t *testing.T) {
	reqBody := `{"value": "test_value"}`
	req, err := http.NewRequest("POST", "/api/encrypt", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EncryptHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"encrypted_value":"encrypted_value_placeholder"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDecryptHandler(t *testing.T) {
	reqBody := `{"value": "test_value"}`
	req, err := http.NewRequest("POST", "/api/decrypt", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DecryptHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"encrypted_value":"decrypted_value_placeholder"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
