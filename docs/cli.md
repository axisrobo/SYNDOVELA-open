# syndovela-cli

`syndovela-cli` is the operator-side companion to the control plane. It
packages bundle manifests and artifacts, benchmarks the resolver and
checks runtime conformance. It is distributed in the `releases/`
directories; verify it before running (see
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
