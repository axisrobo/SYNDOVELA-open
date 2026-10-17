# SYNDOVELA Open

[简体中文](README.zh-CN.md)

> Open SDKs, contracts, examples, and signed binaries for SYNDOVELA, the
> enterprise platform for governing Agent Skill composition across runtimes.

## SYNDOVELA

SYNDOVELA is the control plane for composing and governing reusable Agent
Skills at enterprise scale. It sits between Skill assets and the runtimes
that execute them: Skills are packaged into signable, versioned, composable,
resolvable, distributable, and revocable Bundles and Runtime Profiles.
Runtimes remain responsible for loading, isolation, activation, invocation,
observation, and unloading.

### The Problem

At enterprise scale, treating every reusable Skill as an independently
deployed service turns a code-reuse problem into an operations problem:

```text
10,000 Skills -> 10,000 Services -> 10,000 Containers -> 10,000 Ops Boundaries
```

SYNDOVELA reduces the governance surface without reducing the number of
Skills:

```text
10,000 Skills -> ~1,000 Bundles -> ~100 Runtime Profiles -> a shared runtime substrate
```

### What It Solves

- **Deployment sprawl**: related Skills share a Bundle lifecycle instead of
  each carrying its own deployment boundary.
- **Version drift**: a `ResolutionLock` pins the exact Bundle versions and
  digests a runtime should hold; nothing re-resolves at invocation time.
- **Supply-chain opacity**: every Bundle can carry a signature, SBOM, and
  provenance, while registry presence never implies approval.
- **Unsafe upgrades**: new versions activate before old versions drain; a
  rollback points to a previous lock rather than mutating a release.
- **Ungoverned revocation**: revoking a Bundle or publisher stops new
  invocations across the fleet immediately.

### Open by Construction

SYNDOVELA communicates with runtimes through **SBRP**, the Skill Bundle
Runtime Protocol, an open and vendor-neutral contract.

- Bundles declare `protocol`, `abi`, and `isolation`, never a vendor; the
  same Bundle can run on any conformant runtime.
- Compatibility is based on advertised capabilities, not runtime identity.
- Control planes and runtimes are independently implementable and testable.

See [`docs/sbrp.md`](docs/sbrp.md) and [`docs/conformance.md`](docs/conformance.md).

## SYNDOVELA-open

This Apache-2.0 repository is the public distribution and integration surface
for SYNDOVELA. It contains the open protocol and API contracts, SDKs, examples,
conformance material, and signed release binaries. It contains no SYNDOVELA
core source; the AGPL-3.0-or-later control-plane implementation lives in the
[SYNDOVELA core repository](https://github.com/axisvela/SYNDOVELA).

| Path | Contents |
| --- | --- |
| `docs/sbrp.md` | The Skill Bundle Runtime Protocol |
| `docs/conformance.md` | Conformance checklists for control planes and runtimes |
| `docs/cli.md` | `syndovela-cli` reference (`pack`, `pack-verify`, `warm`, `impact`, `bench`, `conform`) |
| `api/` | Published OpenAPI description of the control-plane surface, mirroring core v1.9.0 |
| `sdk/go/` | Dependency-free Go client and SBRP types |
| `examples/` | Bundle manifests, Runtime Profiles, and runtime descriptors |
| `console/` | Read-only EE dashboard scaffold consuming `/v1/stats` |
| `releases/` | Signed core binaries, checksums, and SBOMs; each release records the core version it ships |
| `scripts/` | Local CI and release checksum/signature verification |

## Local CI

Run the local release checks before publishing:

```powershell
powershell -ExecutionPolicy Bypass -File scripts\ci-local.ps1
```

## Core Concepts

| Object | Meaning |
| --- | --- |
| `Skill` | An invocation unit: one callable operation contract |
| `Bundle` | The packaging, version, and governance unit for related Skills |
| `BundleVersion` | An immutable, digest-addressed release of a Bundle |
| `ResolutionLock` | The deterministic dependency closure of one resolution |
| `RuntimeProfile` | The desired composition of a runtime domain |
| `BundleInstance` | Runtime-authored actual state, reported over SBRP |

## Compatibility

SYNDOVELA follows `major.minor.patch`. The Bundle contract and SBRP are
versioned independently of the product: a breaking change moves to a new
contract major version and the previous one is frozen, never edited. Runtimes
may advertise several SBRP versions at once during migration.

## License

[Apache-2.0](LICENSE). Distributed core binaries carry their own
AGPL-3.0-or-later notice.
