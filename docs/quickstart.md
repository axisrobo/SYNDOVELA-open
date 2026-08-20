# Quickstart

## 1. Install

Download the release for your platform from `releases/`, verify it, and
put it on your path.

```powershell
powershell -ExecutionPolicy Bypass -File scripts\verify-release.ps1 -ReleaseDir releases\v1.0.0\windows-amd64
```

The script checks every binary's SHA-256 against `SHA256SUMS` and, once a
signing key exists, verifies the OpenPGP signature too. See
[`releases/README.md`](../releases/README.md) for the manual commands and
the current signature status.

## 2. Start PostgreSQL and the control plane

```powershell
$env:DATABASE_URL = "postgres://syndovela:syndovela@localhost:5432/syndovela?sslmode=disable"
.\syndovela-migrate.exe
.\syndovela-api.exe
```

Check it is alive:

```powershell
curl http://localhost:8080/healthz
```

## 3. Author a Bundle

Start from `examples/identity-core-bundle.json`. The key decisions:

- Group Skills that share a lifecycle, publisher and isolation
  requirement. Do not create one Bundle per Skill.
- Depend on Skill **contracts**, not on concrete bundle ids, so the
  resolver can choose a provider.
- Declare runtime needs as `protocol`, `abi` and `isolation`. Never name
  a runtime product; doing so makes the bundle unportable for no gain.
- Declare every permission the bundle needs.

## 4. Attach a runtime

SYNDOVELA ships with no built-in runtime. Register an adapter for a
runtime that speaks SBRP �?first-party, third-party or your own. See
[`sbrp.md`](sbrp.md) for what a runtime must implement.

## 5. Register the bundle

```powershell
curl -X POST http://localhost:8080/v1/bundles `
     -H "Content-Type: application/json" `
     -d "@examples/identity-core-bundle.json"
```

Registration puts the bundle in `REGISTERED`. That is not approval. It
still has to pass verification, resolution and approval before it can be
deployed:

```text
REGISTERED -> VERIFIED -> RESOLVED -> APPROVED -> AVAILABLE -> DEPLOYED
```

## 6. Compose a Runtime Profile

Start from `examples/hr-offboarding-profile.json`, then resolve it:

```powershell
curl -X POST http://localhost:8080/v1/runtime-profiles -d "@examples/hr-offboarding-profile.json"
curl -X POST http://localhost:8080/v1/resolutions -d '{"runtimeProfileRef":"hr-offboarding"}'
```

The resolution returns a `ResolutionLock`: pinned versions, digests, the
dependency closure and the resolver inputs. Deployments reference the
lock, never a floating version range.

## 7. Use the Go SDK

```go
package main

import (
	"context"
	"fmt"

	syndovela "github.com/axisrobo/syndovela-open/sdk/go"
)

func main() {
	c := syndovela.New("http://localhost:8080")
	h, err := c.Health(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(h.Product, h.Version)
}
```

## Next

- [`sbrp.md`](sbrp.md) for the runtime protocol, if you are implementing
  or evaluating a runtime.
- [`conformance.md`](conformance.md) for the conformance checklists.
- [`deployment.md`](deployment.md) for topology and scale guidance.
- [`sdk-go.md`](sdk-go.md) for the full client surface.
- [`repository-scope.md`](repository-scope.md) for licensing boundaries.
