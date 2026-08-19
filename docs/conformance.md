# Conformance

Two roles can be certified independently: **bundle control plane** and
**bundle runtime**. Conformance is a property of the implementation, not
of its vendor.

## Runtime conformance

A runtime is conformant when all of the following hold.

### Protocol surface

| # | Requirement |
| --- | --- |
| R1 | Implements `describe`, `fetch`, `validate`, `load`, `activate`, `drain`, `stop`, `unload`, `quarantine`, `report` |
| R2 | `describe` returns a descriptor validating against `sbrp/v1` schema |
| R3 | Advertises only isolation levels it can genuinely enforce |
| R4 | Advertises only ABIs it can genuinely load |

Over-advertising is the most common runtime defect and the most damaging:
the control plane places bundles based on these claims.

### Behaviour

| # | Requirement | How it is tested |
| --- | --- | --- |
| R5 | Re-verifies digest and signature locally before load | Ship an artifact whose digest does not match the binding; load must fail |
| R6 | Never selects versions itself | Provide an ambiguous artifact set; the runtime must load exactly what the lock names |
| R7 | Mediates every invocation, including in-process | Assert a trace and an authorisation check exist for an in-process bundle call |
| R8 | Only `ACTIVE` accepts new invocations | Drive each state and attempt an invocation |
| R9 | Revocation stops new invocations immediately | Revoke during sustained traffic; measure the first refused call |
| R10 | Draining never orphans in-flight invocations | Drain during a long call; the call must complete |
| R11 | Partition does not widen authority | Sever the control plane; the runtime must accept no new bundle |
| R12 | Reports actual state honestly, including failures | Force a load failure; the report must contain `FAILED` |

### Version coexistence

| # | Requirement |
| --- | --- |
| R13 | If `multi-version-coexistence` is advertised, two versions of one bundle can be active simultaneously with correct binding |
| R14 | No state or capability leaks across versions or isolation boundaries |

## Control plane conformance

| # | Requirement | How it is tested |
| --- | --- | --- |
| C1 | Decides eligibility only from advertised capability | Present an unknown runtime with matching capabilities; it must be as eligible as a known one |
| C2 | Emits no product name in the wire format | Inspect the payloads |
| C3 | Produces deterministic locks | Resolve twice from identical inputs; digests must match |
| C4 | Explains every rejection | Force each conflict class; each must yield a code and a reason |
| C5 | Never treats registry presence as approval | Attempt to deploy a merely registered bundle; it must be refused |
| C6 | Rollback targets a previous lock, never a mutated version | Roll back and diff the resulting composition |
| C7 | Ships with no built-in runtime | Start with no adapters registered; no deployment target may exist |

## Self-testing

The reference implementation exercises R8 and C1 as unit tests, which is
the minimum bar rather than the full suite. A complete cross-role test
suite is planned for Wave 3 alongside third-party certification.

Until then, implementers should treat the tables above as the checklist
and publish their results. An implementation that cannot say which rows
it satisfies has not been tested.

## Certification

Certification governance is an enterprise process built on these
primitives. The specification, the schema and the checklists remain
open: certification is about attestation and liability, never about
gating access to the protocol.
