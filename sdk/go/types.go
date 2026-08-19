// Package syndovela is a dependency-free Go client for the SYNDOVELA
// control-plane API.
//
// It is deliberately free of third-party dependencies so that adopting
// it introduces no transitive licensing or supply-chain exposure.
package syndovela

// Bundle is the packaging, version and governance unit for related Skills.
type Bundle struct {
	APIVersion string     `json:"apiVersion"`
	Kind       string     `json:"kind"`
	Metadata   Metadata   `json:"metadata"`
	Skills     []Skill    `json:"skills"`
	Requires   *Requires  `json:"requires,omitempty"`
	Runtime    Runtime    `json:"runtime"`
	Security   Security   `json:"security"`
	Artifacts  []Artifact `json:"artifacts,omitempty"`
}

// Metadata identifies a bundle version.
type Metadata struct {
	ID          string            `json:"id"`
	Version     string            `json:"version"`
	Publisher   string            `json:"publisher"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// Skill is one invocable export of a bundle.
type Skill struct {
	ID             string   `json:"id"`
	Contract       string   `json:"contract"`
	Implementation string   `json:"implementation"`
	CapabilityRefs []string `json:"capabilityRefs,omitempty"`
}

// Requires declares what the bundle needs from its runtime composition.
type Requires struct {
	Skills  []SkillRequirement  `json:"skills,omitempty"`
	Bundles []BundleRequirement `json:"bundles,omitempty"`
}

// SkillRequirement is a contract-first dependency.
type SkillRequirement struct {
	Contract string `json:"contract"`
	Version  string `json:"version"`
}

// BundleRequirement is a hard bundle dependency.
type BundleRequirement struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Reason  string `json:"reason,omitempty"`
}

// Runtime declares execution requirements in vendor-neutral terms.
//
// A bundle states which protocol it speaks, which execution formats it
// ships and which isolation levels it tolerates. It never names a
// runtime product, so the same bundle runs on any conformant runtime.
type Runtime struct {
	// Protocol is the required SBRP version, for example "sbrp/v1".
	Protocol string `json:"protocol"`
	// ABI lists shipped execution formats in order of preference.
	ABI []string `json:"abi"`
	// Isolation lists tolerable boundaries in order of preference.
	Isolation []string `json:"isolation,omitempty"`
	// MinIsolation is a floor, not a preference. Use sparingly: it
	// excludes every runtime that cannot provide the level.
	MinIsolation string `json:"minIsolation,omitempty"`
	// Platforms constrains OS/architecture targets. Empty means portable.
	Platforms []string `json:"platforms,omitempty"`
	// Features lists required optional SBRP features.
	Features []string `json:"features,omitempty"`
	// Constraints carries advisory implementation-specific hints.
	// Resolvers must never gate eligibility on these.
	Constraints map[string]string `json:"constraints,omitempty"`
	Resources   *Resources        `json:"resources,omitempty"`
}

// Resources are advisory limits.
type Resources struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// Security declares supply-chain requirements.
type Security struct {
	Signature   string   `json:"signature"`
	SBOM        string   `json:"sbom"`
	Provenance  string   `json:"provenance"`
	Permissions []string `json:"permissions,omitempty"`
}

// Artifact is a content-addressed payload.
type Artifact struct {
	Digest    string `json:"digest"`
	MediaType string `json:"mediaType"`
	Size      int64  `json:"size"`
	Platform  string `json:"platform,omitempty"`
}

// RuntimeProfile is the desired composition of a runtime domain.
type RuntimeProfile struct {
	Name             string              `json:"name"`
	Generation       int64               `json:"generation"`
	Bundles          []BundleRequirement `json:"bundles"`
	ConfigRefs       []string            `json:"configRefs,omitempty"`
	PolicyRefs       []string            `json:"policyRefs,omitempty"`
	IsolationDefault string              `json:"isolationDefault,omitempty"`
}

// ResolutionLock is the deterministic result of one resolution.
type ResolutionLock struct {
	LockID          string           `json:"lockId"`
	ResolverVersion string           `json:"resolverVersion"`
	Selected        []SelectedBundle `json:"selected"`
	Digest          string           `json:"digest"`
}

// SelectedBundle pins one bundle version by digest.
type SelectedBundle struct {
	BundleID string `json:"bundleId"`
	Version  string `json:"version"`
	Digest   string `json:"digest"`
}
