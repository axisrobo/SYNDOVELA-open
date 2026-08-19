package syndovela

// This file mirrors the Skill Bundle Runtime Protocol (SBRP) types.
//
// SBRP is an open, vendor-neutral protocol between a bundle control
// plane and a bundle runtime. These types are published here so that
// runtime authors can implement the protocol without depending on any
// control-plane implementation.
//
// Implementing SBRP requires no relationship with AxisRobo and no use of
// AGPL source. The specification lives at contracts/sbrp/v1 in the core
// repository.

import "time"

// ProtocolVersion is the SBRP version these types describe.
const ProtocolVersion = "sbrp/v1"

// RuntimeDescriptor is a runtime's self-description.
//
// Control planes decide eligibility from these fields alone.
// Implementation and ImplementationVersion are informational; a control
// plane that branches on them has broken protocol openness.
type RuntimeDescriptor struct {
	RuntimeID             string            `json:"runtimeId"`
	Implementation        string            `json:"implementation,omitempty"`
	ImplementationVersion string            `json:"implementationVersion,omitempty"`
	ProtocolVersions      []string          `json:"protocolVersions"`
	Isolation             []string          `json:"isolation"`
	ABIs                  []string          `json:"abis"`
	Platform              string            `json:"platform"`
	Features              []string          `json:"features,omitempty"`
	Limits                map[string]string `json:"limits,omitempty"`
	Labels                map[string]string `json:"labels,omitempty"`
}

// Optional SBRP features a runtime may advertise.
const (
	FeatureHotSwap       = "hot-swap"
	FeatureOfflinePack   = "offline-pack"
	FeatureAttestation   = "attestation"
	FeatureResourceQuota = "resource-quota"
	FeatureMultiVersion  = "multi-version-coexistence"
)

// BundleBinding is one bundle version a runtime is asked to host.
type BundleBinding struct {
	BundleID          string   `json:"bundleId"`
	Version           string   `json:"version"`
	Digest            string   `json:"digest"`
	ArtifactURIs      []string `json:"artifactUris,omitempty"`
	ResolutionLockRef string   `json:"resolutionLockRef"`
	CompositionRef    string   `json:"compositionRef,omitempty"`
	IsolationRequired string   `json:"isolationRequired,omitempty"`
	ABI               string   `json:"abi,omitempty"`
	PolicyRefs        []string `json:"policyRefs,omitempty"`
}

// InstanceState is the runtime-side lifecycle state.
type InstanceState string

// SBRP instance states. Only ACTIVE accepts new invocations; every other
// state is fail-closed.
const (
	InstanceFetched     InstanceState = "FETCHED"
	InstanceValidated   InstanceState = "VALIDATED"
	InstanceLoaded      InstanceState = "LOADED"
	InstanceActive      InstanceState = "ACTIVE"
	InstanceDraining    InstanceState = "DRAINING"
	InstanceStopped     InstanceState = "STOPPED"
	InstanceUnloaded    InstanceState = "UNLOADED"
	InstanceFailed      InstanceState = "FAILED"
	InstanceQuarantined InstanceState = "QUARANTINED"
)

// AcceptsNewInvocations reports whether an instance in this state may
// receive new skill invocations.
func AcceptsNewInvocations(s InstanceState) bool {
	return s == InstanceActive
}

// BundleInstance is runtime-authored actual state. A control plane never
// authors it, and a runtime that reports desired state as actual state
// is non-conformant.
type BundleInstance struct {
	InstanceID        string        `json:"instanceId"`
	RuntimeID         string        `json:"runtimeId"`
	NodeID            string        `json:"nodeId,omitempty"`
	BundleID          string        `json:"bundleId"`
	Version           string        `json:"version"`
	Digest            string        `json:"digest"`
	State             InstanceState `json:"state"`
	Isolation         string        `json:"isolation"`
	ActiveInvocations int64         `json:"activeInvocations"`
	Health            string        `json:"health,omitempty"`
	LoadedAt          time.Time     `json:"loadedAt"`
	ReportedAt        time.Time     `json:"reportedAt"`
}

// SkillInvocation is the mediated call record. Mediation is required
// even for bundles running in the caller's own process, so that
// authorisation, timeout, cancellation, tracing and resource accounting
// are preserved.
type SkillInvocation struct {
	InvocationID   string   `json:"invocationId"`
	Caller         string   `json:"caller"`
	InstanceID     string   `json:"instanceId"`
	SkillContract  string   `json:"skillContract"`
	Operation      string   `json:"operation"`
	GrantRef       string   `json:"grantRef,omitempty"`
	IdempotencyKey string   `json:"idempotencyKey,omitempty"`
	EffectRefs     []string `json:"effectRefs,omitempty"`
	TraceID        string   `json:"traceId,omitempty"`
}

// ActualStateReport is the reconciliation acknowledgement a runtime
// sends back.
type ActualStateReport struct {
	TargetRef         string           `json:"targetRef"`
	RuntimeID         string           `json:"runtimeId"`
	DesiredGeneration int64            `json:"desiredGeneration"`
	ActualGeneration  int64            `json:"actualGeneration"`
	Instances         []BundleInstance `json:"instances"`
	DriftReason       string           `json:"driftReason,omitempty"`
	ObservedAt        time.Time        `json:"observedAt"`
}
