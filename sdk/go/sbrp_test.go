package syndovela

import "testing"

func TestOnlyActiveAcceptsInvocations(t *testing.T) {
	states := []InstanceState{
		InstanceFetched, InstanceValidated, InstanceLoaded,
		InstanceDraining, InstanceStopped, InstanceUnloaded,
		InstanceFailed, InstanceQuarantined,
	}
	for _, s := range states {
		if AcceptsNewInvocations(s) {
			t.Fatalf("state %s must not accept new invocations", s)
		}
	}
	if !AcceptsNewInvocations(InstanceActive) {
		t.Fatal("ACTIVE must accept new invocations")
	}
}

// A bundle manifest must be expressible without naming a runtime
// product, otherwise the format is not portable.
func TestRuntimeRequirementsAreVendorNeutral(t *testing.T) {
	r := Runtime{
		Protocol:  ProtocolVersion,
		ABI:       []string{"wasi/preview2"},
		Isolation: []string{"wasm", "process"},
	}
	if r.Protocol != "sbrp/v1" {
		t.Fatalf("got protocol %q, want sbrp/v1", r.Protocol)
	}
	if len(r.ABI) == 0 {
		t.Fatal("a bundle must declare at least one ABI")
	}
}
