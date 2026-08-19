package syndovela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client talks to a SYNDOVELA control plane.
type Client struct {
	baseURL string
	http    *http.Client
}

// Option customises a Client.
type Option func(*Client)

// WithHTTPClient overrides the underlying HTTP client.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.http = h }
}

// New builds a client against the given control-plane base URL.
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		http:    &http.Client{Timeout: 30 * time.Second},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// Health is the control-plane liveness payload.
type Health struct {
	Product string    `json:"product"`
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
}

// Health reports control-plane liveness and version.
func (c *Client) Health(ctx context.Context) (Health, error) {
	var out Health
	err := c.do(ctx, http.MethodGet, "/healthz", nil, &out)
	return out, err
}

// RegisterBundle registers a bundle manifest with the control plane.
func (c *Client) RegisterBundle(ctx context.Context, b Bundle) (Bundle, error) {
	var out Bundle
	err := c.do(ctx, http.MethodPost, "/v1/bundles", b, &out)
	return out, err
}

// Resolve turns requirements into a deterministic ResolutionLock.
func (c *Client) Resolve(ctx context.Context, req any) (ResolutionLock, error) {
	var out ResolutionLock
	err := c.do(ctx, http.MethodPost, "/v1/resolutions", req, &out)
	return out, err
}

// ApplyRuntimeProfile creates or updates a desired runtime composition.
func (c *Client) ApplyRuntimeProfile(ctx context.Context, p RuntimeProfile) (RuntimeProfile, error) {
	var out RuntimeProfile
	err := c.do(ctx, http.MethodPost, "/v1/runtime-profiles", p, &out)
	return out, err
}

// APIError is a non-2xx response from the control plane.
type APIError struct {
	StatusCode int
	Body       string
}

// Error implements error.
func (e *APIError) Error() string {
	return fmt.Sprintf("syndovela: status %d: %s", e.StatusCode, e.Body)
}

func (c *Client) do(ctx context.Context, method, path string, in, out any) error {
	var body io.Reader
	if in != nil {
		raw, err := json.Marshal(in)
		if err != nil {
			return fmt.Errorf("syndovela: encode request: %w", err)
		}
		body = bytes.NewReader(raw)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return fmt.Errorf("syndovela: build request: %w", err)
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("syndovela: %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("syndovela: read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{StatusCode: resp.StatusCode, Body: string(raw)}
	}
	if out == nil || len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("syndovela: decode response: %w", err)
	}
	return nil
}
