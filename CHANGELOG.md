# Changelog

All notable changes to SYNDOVELA Open are documented here. This project
follows Keep a Changelog and semantic versioning; patch versions carry
backward-compatible fixes and documentation only.

## [1.2.0] - 2026

### Added
- **Core v1.8.0 binaries.** All 15 binaries (5 commands × 3 platforms)
  rebuilt from core `48be35a7d5` (v1.8.0), closing the 7-minor-version
  gap: released binaries now include tenant rate limiting, bearer-token
  auth + RBAC, lifecycle webhooks, the event store, change sets, the
  SBOM endpoint and the bench baseline.
- **Tag mapping table.** `releases/README.md` now records which Core
  version each Open release ships and states explicitly that Open and
  Core tag numbers are independent semver streams.

## [1.1.0] - 2026

### Added
- **Real release signing.** v1.1.0 binaries are released under a
  provisioned ed25519 signing key (`884C89AA 4F04D13A 716F6DA4 1D0E3681
  55C1F509`). Every platform's `SHA256SUMS` and the release `sbom.json`
  are detached-signed (`SHA256SUMS.asc`, `sbom.json.asc`), and the
  armored public key is committed at
  `releases/SYNDOVELA-RELEASE-SIGNING-KEY.asc`.
- `sbom.json` (SPDX-2.3) covering all 15 binaries across the three
  platforms, with per-binary SHA-256 digests.
- `verify-release` / `verify-release.ps1` now import the committed
  release key when it is absent from the local keyring, so verification
  works on a machine that has never seen the key.

### Changed
- v1.0.0 signatures remain placeholders (pre-key); v1.1.0 onwards the
  OpenPGP signature is authoritative alongside the checksums.

## [1.0.2] - 2026

### Fixed
- `scripts/verify-release` and `scripts/verify-release.ps1` now take the
  release **directory** to verify (`-ReleaseDir` for PowerShell), matching the
  `releases/README.md` and `docs/quickstart.md` usage.
- The POSIX script tolerates CRLF line endings and a UTF-8 BOM in
  `SHA256SUMS`, so checksum verification works on files checked out on
  Windows.
- `releases/README.md` and `docs/quickstart.md` examples updated to the
  actual `-ReleaseDir` interface.

## [1.0.1] - 2026

### Added
- `scripts/verify-release` (POSIX sh) and `scripts/verify-release.ps1`
  (PowerShell) for verifying a release directory: every binary listed in
  `SHA256SUMS` is checked for existence and matching SHA-256, and a real
  (non-placeholder) OpenPGP signature is additionally verified with gpg.
- `docs/cli.md` documenting the `syndovela-cli` `pack`, `bench` and
  `conform` commands.

### Changed
- `releases/README.md` and `docs/quickstart.md` now point at the
  verification script. The manual GPG/`Get-FileHash` commands remain for
  completeness.
- Signature status is now explicit: the v1.0.0 `SHA256SUMS.asc` files are
  placeholders (no signing key provisioned), so checksums are the
  authoritative verification and the script downgrades a placeholder
  signature to a warning.

## [1.0.0] - 2026

### Added
- Initial release: core binaries for windows-amd64, linux-amd64 and
  darwin-arm64 with checksums and notices; Go and TypeScript SDKs; the
  SBRP specification and independent-review checklist.
