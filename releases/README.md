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

```powershell
gpg --import SYNDOVELA-RELEASE-SIGNING-KEY.asc
gpg --verify SHA256SUMS.asc SHA256SUMS
Get-FileHash .\syndovela-api.exe -Algorithm SHA256
```

Verify the signature before trusting the checksums, and verify the
checksums before running the binaries.

## Licensing note

The binaries published here are AGPL-3.0-or-later. This repository's
Apache-2.0 license covers the SDK, examples, API description and
documentation, not the distributed core binaries.

## Index

No releases published yet. v0.1.0 is the repository baseline.
