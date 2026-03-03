package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

// requestLog captures the last request query parameters per path.
type requestLog struct {
	mu   sync.Mutex
	last map[string]url.Values
}

func (rl *requestLog) params(path string) url.Values {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.last[path]
}

func (rl *requestLog) assertParam(t *testing.T, path, key, want string) {
	t.Helper()
	got := rl.params(path).Get(key)
	if got != want {
		t.Errorf("expected %s=%q for %s, got %q", key, want, path, got)
	}
}

func (rl *requestLog) assertNoParam(t *testing.T, path, key string) {
	t.Helper()
	if got := rl.params(path).Get(key); got != "" {
		t.Errorf("expected no %s param for %s, got %q", key, path, got)
	}
}

// newTestServer creates a mock HTTP server that returns the given response for matching paths
// and captures request query parameters for assertion.
func newTestServer(t *testing.T, handlers map[string]any) (*requestLog, *Client) {
	t.Helper()
	rl := &requestLog{last: make(map[string]url.Values)}
	mux := http.NewServeMux()
	for path, resp := range handlers {
		body, err := json.Marshal(resp)
		if err != nil {
			t.Fatalf("marshaling mock response: %v", err)
		}
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			rl.mu.Lock()
			rl.last[path] = r.URL.Query()
			rl.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		})
	}
	server := httptest.NewServer(mux)
	client := NewWithBaseURL("test-key", server.URL, "test")
	t.Cleanup(server.Close)
	return rl, client
}

func TestAPIErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "unauthorized"}`))
	}))
	defer server.Close()

	c := NewWithBaseURL("bad-key", server.URL, "test")
	_, err := c.GetInstruments()
	if err == nil {
		t.Fatal("expected error for unauthorized request")
	}
}
