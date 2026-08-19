# Quickstart

## 1. Install

Download the release for your platform from `releases/`, verify it, and
put it on your path.

```powershell
# verify before running
Get-FileHash .\syndovela-api.exe -Algorithm SHA256
# compare against SHA256SUMS, then verify SHA256SUMS.asc with the
# SYNDOVELA release signing key
```

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
- Declare every permission the bundle needs.

## 4. Register it

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

## 5. Compose a Runtime Profile

Start from `examples/hr-offboarding-profile.json`, then resolve it:

```powershell
curl -X POST http://localhost:8080/v1/runtime-profiles -d "@examples/hr-offboarding-profile.json"
curl -X POST http://localhost:8080/v1/resolutions -d '{"runtimeProfileRef":"hr-offboarding"}'
```

The resolution returns a `ResolutionLock`: pinned versions, digests, the
dependency closure and the resolver inputs. Deployments reference the
lock, never a floating version range.

## 6. Use the Go SDK

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

- [`deployment.md`](deployment.md) for topology and scale guidance.
- [`sdk-go.md`](sdk-go.md) for the full client surface.
- [`repository-scope.md`](repository-scope.md) for licensing boundaries.
