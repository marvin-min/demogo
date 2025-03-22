package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	req := httptest.NewRequest("GET", "/World", nil)
	w := httptest.NewRecorder()
	HelloServer(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}

	expectedBody := "Hello, World!"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q; got %q", expectedBody, w.Body.String())
	}
}
