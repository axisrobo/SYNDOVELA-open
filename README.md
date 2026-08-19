# SYNDOVELA Open

> Enterprise Skill Bundle & Runtime Composition Platform.

SYNDOVELA turns large populations of reusable Agent Skills into a much
smaller number of versioned, signed and resolvable Bundles and Runtime
Profiles, then continuously reconciles those compositions with the
runtimes that execute them.

## The Problem

Once an enterprise has thousands of Agent Skills, treating each Skill as
its own service turns a code-reuse problem into an operations problem:

```text
10,000 Skills -> 10,000 Services -> 10,000 Containers -> 10,000 Ops Boundaries
```

SYNDOVELA compresses the governance surface instead of the Skill count:

```text
10,000 Skills -> ~1,000 Bundles -> ~100 Runtime Profiles -> a shared runtime substrate
```

## What It Solves

- **Deployment sprawl** — related Skills share one Bundle lifecycle
  rather than each carrying its own deployment boundary.
- **Version drift** — a `ResolutionLock` pins the exact bundle versions
  and digests a runtime should hold, so nothing re-resolves at call time.
- **Supply-chain opacity** — every Bundle carries a signature, SBOM and
  provenance, and registry presence never implies approval.
- **Unsafe upgrades** — new versions activate before old ones drain, and
  rollback means pointing at a previous lock, not mutating a release.
- **Ungoverned revocation** — revoking a Bundle or publisher stops new
  invocations across the fleet immediately.

## Concepts

| Object | Meaning |
| --- | --- |
| `Skill` | An invocation unit: one callable operation contract |
| `Bundle` | The packaging, version and governance unit for related Skills |
| `BundleVersion` | An immutable, digest-addressed release of a Bundle |
| `ResolutionLock` | The deterministic dependency closure of one resolution |
| `RuntimeProfile` | The desired composition of a runtime domain |

## This Repository

| Path | Contents |
| --- | --- |
| `api/` | Published OpenAPI description of the control-plane surface |
| `sdk/go/` | Dependency-free Go client |
| `examples/` | Example Bundle manifests and Runtime Profiles |
| `docs/` | Quickstart, deployment and repository scope |
| `releases/` | Signed core binaries, checksums and SBOMs |

This repository contains no core source. The SYNDOVELA control-plane
implementation is licensed under AGPL-3.0-or-later and lives in a
separate repository.

## Compatibility

SYNDOVELA follows `major.minor.patch`. The published Bundle contract is
versioned independently; a breaking change moves to a new contract major
version and the previous one is frozen, never edited.

## License

[Apache-2.0](LICENSE). Distributed core binaries carry their own
AGPL-3.0-or-later notice.
