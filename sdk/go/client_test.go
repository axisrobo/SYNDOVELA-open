package syndovela

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/healthz" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"product":"syndovela","version":"0.1.0"}`))
	}))
	defer srv.Close()

	got, err := New(srv.URL).Health(context.Background())
	if err != nil {
		t.Fatalf("Health: %v", err)
	}
	if got.Version != "0.1.0" {
		t.Fatalf("got version %q, want 0.1.0", got.Version)
	}
}

func TestAPIErrorOnNon2xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte(`{"code":"NOT_IMPLEMENTED"}`))
	}))
	defer srv.Close()

	_, err := New(srv.URL).RegisterBundle(context.Background(), Bundle{})
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("got %T, want *APIError", err)
	}
	if apiErr.StatusCode != http.StatusNotImplemented {
		t.Fatalf("got status %d, want 501", apiErr.StatusCode)
	}
}

func TestBaseURLTrailingSlashTrimmed(t *testing.T) {
	if got := New("https://example.com/").baseURL; got != "https://example.com" {
		t.Fatalf("got %q, want https://example.com", got)
	}
}
