# syndovela-cli

`syndovela-cli` is the operator-side companion to the control plane. It
packages bundle manifests and artifacts, verifies offline packs, warms
and measures the resolver cache, fetches impact reports, benchmarks the
resolver and checks runtime conformance. It is distributed in the
`releases/` directories; verify it before running (see
[`verify-release`](../scripts/verify-release)).

```powershell
syndovela-cli <command> [flags]
```

## pack

`pack` builds an offline bundle set from a manifest and an artifact file.
The artifact's SHA-256 must equal the digest declared in the manifest, so
a pack is only ever produced from content that matches its published
addresses.

```powershell
syndovela-cli pack --manifest identity-core.json --artifact identity-core.wasm --out pack.json
```

| Flag | Meaning |
| --- | --- |
| `--manifest` (required) | Bundle manifest JSON. The pack key is `metadata.id@metadata.version`. |
| `--artifact` | Artifact file; its SHA-256 must equal `artifacts[0].digest`. |
| `--out` | Output path (default `pack.json`). |

The resulting pack bundles the manifest and any matching artifacts so it
can be loaded into a registry without re-fetching content from a remote
feed.

## pack-verify

`pack-verify` checks an offline pack's internal digest consistency and,
given `--lock`, that the pack is exactly the composition the lock names.
It is the offline counterpart of the control plane's verification gate.

```powershell
syndovela-cli pack-verify --pack pack.json --lock resolution-lock.json
```

| Flag | Meaning |
| --- | --- |
| `--pack` (required) | Pack file produced by `pack`. |
| `--lock` | ResolutionLock to check the pack against (`distribution.VerifyOffline`). |

## warm

`warm` resolves a synthetic catalog through the resolver cache twice and
reports the second-pass hit rate — the number to quote when sizing the
distributed cache.

```powershell
syndovela-cli warm --bundles 1000 --skills 10
```

## impact

`impact` fetches the dependency/deployment impact report for a bundle
version from a running control plane and writes it to stdout or a file.

```powershell
syndovela-cli impact --server http://localhost:8080 identity-core 1.0.0 --out impact.json
```

| Flag | Meaning |
| --- | --- |
| `--server` (required) | Control-plane base URL. |
| `--out` | Write the report here instead of stdout. |

## bench

`bench` runs a synthetic resolution benchmark through the same resolver
the control plane uses, reporting throughput and latency. It is the
repeatable number you should run before and after any resolver change.

```powershell
syndovela-cli bench --bundles 1000 --skills 10 --iterations 100
```

| Flag | Meaning | Default |
| --- | --- | --- |
| `--bundles` | Synthetic catalog size | 1000 |
| `--skills` | Skills per bundle | 10 |
| `--iterations` | Resolutions to run | 100 |

Output includes the experiment id, resolved/conflict counts and the
p95 resolution latency in milliseconds. Use a fixed flag set when
comparing runs so the numbers stay comparable.

## conform

`conform` runs the SBRP conformance suite against a runtime's HTTP
transport and reports each check PASS/FAIL.

```powershell
syndovela-cli conform --runtime http://localhost:9000
```

| Flag | Meaning |
| --- | --- |
| `--runtime` (required) | Base URL of an SBRP HTTP runtime. |

See [`sbrp.md`](sbrp.md) and the
[independent-review checklist](sbrp-review-checklist.md) for what the
checks cover.
