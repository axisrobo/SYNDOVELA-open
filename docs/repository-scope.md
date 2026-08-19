# Repository Scope

SYNDOVELA is distributed across three repositories with different
licenses and different audiences.

| Repository | Purpose | License |
| --- | --- | --- |
| `syndovela-open` (this repo) | SDK, examples, published API, docs, signed release binaries | Apache-2.0 |
| `syndovela` | Control-plane implementation | AGPL-3.0-or-later |
| `syndovela-ee` | Enterprise features | Commercial |

## What is in this repository

- The published OpenAPI description of the control-plane surface.
- A dependency-free Go SDK.
- Example Bundle manifests and Runtime Profiles.
- Quickstart and deployment documentation.
- Signed release binaries with checksums, SBOMs and license notices.

## What is not in this repository

- Core control-plane source code. It is AGPL-3.0-or-later and lives in
  the `syndovela` repository. Nothing here is AGPL source.
- Enterprise features: multi-tenant registry, approval workflows, fleet
  reconciliation, federation, compliance packs.

## Using the SDK

The Go SDK has zero third-party dependencies. Adopting it introduces no
transitive licensing or supply-chain exposure. It is Apache-2.0, so it
may be linked into proprietary software freely.

Running the core control-plane binaries is a different matter: those
binaries are AGPL-3.0-or-later and each release directory ships the
corresponding notice.
