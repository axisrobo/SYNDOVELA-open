# Releases

Signed SYNDOVELA Core release binaries.

Each `vX.Y.Z/` directory contains, per platform:

| File | Purpose |
| --- | --- |
| `syndovela-api`, `syndovela-migrate`, `syndovela-mcp` | Binaries |
| `SHA256SUMS` | Checksums for that platform's binaries |
| `SHA256SUMS.asc` | Detached signature over `SHA256SUMS` |
| `LICENSE-SYNDOVELA-CORE.txt` | AGPL-3.0-or-later notice for the binaries |
| `NOTICE` | Originating core commit and target platform |

Each release directory also contains `sbom.json` and `sbom.json.asc`,
plus the release signing public key.

## Verifying a release

Use the verification script, which checks every binary's SHA-256 against
`SHA256SUMS` and additionally verifies the OpenPGP signature when one is
present. Pass the release directory to verify:

```powershell
powershell -ExecutionPolicy Bypass -File scripts\verify-release.ps1 -ReleaseDir releases\v1.2.0\windows-amd64
```

```sh
./scripts/verify-release releases/v1.2.0/windows-amd64
```

The script exits `0` on success, `1` if any checksum check fails, and `2`
if the signature check was skipped or degraded.

Manual equivalent:

```powershell
gpg --import SYNDOVELA-RELEASE-SIGNING-KEY.asc
gpg --verify SHA256SUMS.asc SHA256SUMS
Get-FileHash .\syndovela-api.exe -Algorithm SHA256
```

> **Signature status:** starting with v1.1.0 every `SHA256SUMS` and
> `sbom.json` is signed with the SYNDOVELA release key, so both the
> checksums and the OpenPGP signature are authoritative. The release
> public key is committed at `releases/SYNDOVELA-RELEASE-SIGNING-KEY.asc`:
>
> - Fingerprint: `884C89AA 4F04D13A 716F6DA4 1D0E3681 55C1F509`
> - UID: `SYNDOVELA Release Signing <releases@syndovela.dev>`
>
> The v1.0.0 release predates the key and its `SHA256SUMS.asc` files are
> placeholders; verify those checksums manually.

## Licensing note

The binaries published here are AGPL-3.0-or-later. This repository's
Apache-2.0 license covers the SDK, examples, API description and
documentation, not the distributed core binaries.

## Index

| Version | Platforms | Binaries | Core version shipped |
| --- | --- | --- | --- |
| v1.2.0 | windows-amd64, linux-amd64, darwin-arm64 | syndovela-api, syndovela-cli, syndovela-mcp, syndovela-migrate, syndovela-rt | **core v1.8.0** (`48be35a7d5`) |
| v1.1.0 | windows-amd64, linux-amd64, darwin-arm64 | syndovela-api, syndovela-cli, syndovela-mcp, syndovela-migrate, syndovela-rt | core v1.1.0 (`3ede36d29539`) |
| v1.0.0 | windows-amd64, linux-amd64, darwin-arm64 | syndovela-api, syndovela-cli, syndovela-mcp, syndovela-migrate, syndovela-rt | core v1.0.x |

> **Tag mapping:** Open release tags and Core source tags are
> independent semver streams — `open vX.Y.Z` packages the Core version
> listed above (authoritative record: each platform's `NOTICE` file).
> Do not assume equal numbers mean equal content; always check the
> "Core version shipped" column or the `NOTICE`.