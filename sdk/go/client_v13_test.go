package syndovela

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestV13Methods(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/bundles", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`[{"id":"a","version":"1.0.0"}]`))
	})
	mux.HandleFunc("GET /v1/bundles/a/1.0.0/sbom", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"spdxVersion":"SPDX-2.3","name":"a@1.0.0"}`))
	})
	mux.HandleFunc("POST /v1/change-sets", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"id":"s1","states":["VERIFIED"]}`))
	})
	mux.HandleFunc("GET /v1/audit", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "limit=5" {
			t.Errorf("audit query %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{"entries":[{"actor":"op","action":"bundle.register"}]}`))
	})
	mux.HandleFunc("GET /v1/events", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("aggregate") != "a@1.0.0" {
			t.Errorf("events aggregate %q", r.URL.Query().Get("aggregate"))
		}
		_, _ = w.Write([]byte(`{"aggregate":"a@1.0.0","events":[{"id":1,"type":"registered"}]}`))
	})
	mux.HandleFunc("GET /v1/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		_, _ = w.Write([]byte("syndovela_uptime_seconds 42\n"))
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()
	c := New(srv.URL)
	ctx := context.Background()

	bundles, err := c.ListBundles(ctx)
	if err != nil || len(bundles) != 1 {
		t.Fatalf("ListBundles: %v %v", err, bundles)
	}

	sbom, err := c.SBOM(ctx, "a", "1.0.0")
	if err != nil || sbom["name"] != "a@1.0.0" {
		t.Fatalf("SBOM: %v %v", err, sbom)
	}

	states, err := c.ApplyChangeSet(ctx, ChangeSet{ID: "s1"})
	if err != nil || states["id"] != "s1" {
		t.Fatalf("ApplyChangeSet: %v %v", err, states)
	}

	entries, err := c.Audit(ctx, 5)
	if err != nil || len(entries) != 1 || entries[0].Action != "bundle.register" {
		t.Fatalf("Audit: %v %v", err, entries)
	}

	events, err := c.Events(ctx, "a@1.0.0")
	if err != nil || len(events) != 1 || events[0].Type != "registered" {
		t.Fatalf("Events: %v %v", err, events)
	}

	raw, err := c.MetricsRaw(ctx)
	if err != nil || !strings.Contains(raw, "syndovela_uptime_seconds 42") {
		t.Fatalf("MetricsRaw: %v %q", err, raw)
	}
}
