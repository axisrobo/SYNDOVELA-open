package syndovela

import (
	"context"
	"net/http"
)

// GetBundle fetches one recorded bundle version.
func (c *Client) GetBundle(ctx context.Context, id, version string) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodGet, "/v1/bundles/"+id+"/"+version, nil, &out)
	return out, err
}

// ListBundles lists every recorded bundle version.
func (c *Client) ListBundles(ctx context.Context) ([]map[string]any, error) {
	var out []map[string]any
	err := c.do(ctx, http.MethodGet, "/v1/bundles", nil, &out)
	return out, err
}

// BundleAction runs an action (:verify / :revoke) on a bundle version.
func (c *Client) BundleAction(ctx context.Context, id, version, action string) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodPost, "/v1/bundles/"+id+"/"+version+":"+action, nil, &out)
	return out, err
}

// Impact fetches the dependency/deployment impact report for a bundle
// version.
func (c *Client) Impact(ctx context.Context, id, version string) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodGet, "/v1/bundles/"+id+"/"+version+"/impact", nil, &out)
	return out, err
}

// SBOM fetches the SPDX-lite SBOM document for a bundle version.
func (c *Client) SBOM(ctx context.Context, id, version string) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodGet, "/v1/bundles/"+id+"/"+version+"/sbom", nil, &out)
	return out, err
}

// ChangeSet is one atomic set of bundle state transitions.
type ChangeSet struct {
	ID      string                   `json:"id"`
	Changes []map[string]interface{} `json:"changes"`
}

// ApplyChangeSet applies an atomic change set and returns the resulting
// states.
func (c *Client) ApplyChangeSet(ctx context.Context, set ChangeSet) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodPost, "/v1/change-sets", set, &out)
	return out, err
}

// AuditEntry is one audit record.
type AuditEntry struct {
	TenantID   string `json:"tenantId"`
	Actor      string `json:"actor"`
	Action     string `json:"action"`
	SubjectRef string `json:"subjectRef"`
	Detail     string `json:"detail"`
	OccurredAt string `json:"occurredAt"`
}

// Audit lists recent audit entries for the tenant.
func (c *Client) Audit(ctx context.Context, limit int) ([]AuditEntry, error) {
	path := "/v1/audit"
	if limit > 0 {
		path += "?limit=" + itoa(limit)
	}
	var out struct {
		Entries []AuditEntry `json:"entries"`
	}
	err := c.do(ctx, http.MethodGet, path, nil, &out)
	return out.Entries, err
}

// BenchBaseline is the checked-in SYNDO-Bench baseline document.
func (c *Client) BenchBaseline(ctx context.Context) (map[string]any, error) {
	var out map[string]any
	err := c.do(ctx, http.MethodGet, "/v1/bench/baseline", nil, &out)
	return out, err
}

// Event is one lifecycle event from the append-only store.
type Event struct {
	ID        int64  `json:"id"`
	Aggregate string `json:"aggregate"`
	Type      string `json:"type"`
	Seq       int64  `json:"seq"`
	At        string `json:"at"`
}

// Events replays lifecycle events for one aggregate (bundleID@version).
func (c *Client) Events(ctx context.Context, aggregate string) ([]Event, error) {
	var out struct {
		Events []Event `json:"events"`
	}
	err := c.do(ctx, http.MethodGet, "/v1/events?aggregate="+aggregate, nil, &out)
	return out.Events, err
}

// MetricsRaw returns the Prometheus text exposition from /v1/metrics.
func (c *Client) MetricsRaw(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/v1/metrics", nil)
	if err != nil {
		return "", err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", &APIError{StatusCode: resp.StatusCode}
	}
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 2048)
	for {
		n, err := resp.Body.Read(tmp)
		buf = append(buf, tmp[:n]...)
		if err != nil {
			break
		}
	}
	return string(buf), nil
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var digits [20]byte
	i := len(digits)
	for n > 0 {
		i--
		digits[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		digits[i] = '-'
	}
	return string(digits[i:])
}
