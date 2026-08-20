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
powershell -ExecutionPolicy Bypass -File scripts\verify-release.ps1 -ReleaseDir releases\v1.0.0\windows-amd64
```

```sh
./scripts/verify-release releases/v1.0.0/windows-amd64
```

The script exits `0` on success, `1` if any checksum check fails, and `2`
if the signature check was skipped or degraded (e.g. a placeholder
signature).

Manual equivalent:

```powershell
gpg --import SYNDOVELA-RELEASE-SIGNING-KEY.asc
gpg --verify SHA256SUMS.asc SHA256SUMS
Get-FileHash .\syndovela-api.exe -Algorithm SHA256
```

> **Signature status:** as of v1.0.0 the `SHA256SUMS.asc` files are
> placeholders — no GPG key has been provisioned yet, so the checksums
> are the authoritative verification. The script reports a placeholder
> signature as a warning rather than a failure. Verify the checksums
> before running the binaries, and treat the placeholders as a known
> limitation until the release key lands.

## Licensing note

The binaries published here are AGPL-3.0-or-later. This repository's
Apache-2.0 license covers the SDK, examples, API description and
documentation, not the distributed core binaries.

## Index

| Version | Platforms | Binaries |
| --- | --- | --- |
| v1.0.0 | windows-amd64, linux-amd64, darwin-arm64 | syndovela-api, syndovela-cli, syndovela-mcp, syndovela-migrate, syndovela-rt |