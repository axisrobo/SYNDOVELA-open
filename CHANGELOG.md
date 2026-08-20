# Changelog

All notable changes to SYNDOVELA Open are documented here. This project
follows Keep a Changelog and semantic versioning; patch versions carry
backward-compatible fixes and documentation only.

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
