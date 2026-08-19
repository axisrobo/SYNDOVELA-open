# Skill Bundle Runtime Protocol (SBRP)

SBRP is an open, vendor-neutral protocol between a **bundle control
plane** and a **bundle runtime**.

It exists so the Bundle format and the runtimes that execute it evolve
independently. Implementing either side requires no relationship with
AxisRobo and no use of AGPL-licensed source: the specification, the JSON
Schema and the Go types in this repository are Apache-2.0.

## Why a protocol rather than an integration

A control plane that talks to one runtime is a product feature. A control
plane that talks to any conformant runtime is infrastructure. The
difference is enforced by two rules:

1. No product name appears in the wire format.
2. Compatibility is decided from advertised capability, never from
   runtime identity.

A control plane that special-cases a particular runtime, or a bundle
manifest that names one, has broken the guarantee even if it still
technically parses.

## Roles

| Role | Owns | Must never |
| --- | --- | --- |
| Control plane | Desired composition, registry, resolution, approval, revocation intent | Load code, run an agent loop, hold effect or checkpoint state |
| Runtime | Actual state, isolation, activation, invocation mediation, resource accounting | Re-resolve dependencies, fetch unapproved bundles, trust control-plane integrity claims without local verification |

The asymmetry is deliberate: the control plane knows what *should* run,
and only the runtime knows what *is* running.

## Capability negotiation

A runtime describes itself. See
[`../examples/runtime-descriptor.json`](../examples/runtime-descriptor.json).

```json
{
  "runtimeId": "node-a7",
  "implementation": "example-wasm-host",
  "protocolVersions": ["sbrp/v1"],
  "isolation": ["wasm", "process"],
  "abis": ["wasi/preview2"],
  "platform": "linux/amd64",
  "features": ["hot-swap"]
}
```

`implementation` and `implementationVersion` exist for operators and
logs. They are not compatibility gates.

Correspondingly, a bundle declares needs, not vendors:

```json
{
  "runtime": {
    "protocol": "sbrp/v1",
    "abi": ["wasi/preview2", "native/grpc"],
    "isolation": ["wasm", "process", "container"],
    "platforms": ["linux/amd64"]
  }
}
```

`isolation` is a preference list, `minIsolation` is a floor. Use the
floor sparingly: it excludes every runtime that cannot provide the level.

## Operations

| Verb | Direction | Meaning |
| --- | --- | --- |
| `describe` | plane → runtime | Negotiate capabilities before shipping anything |
| `fetch` | plane → runtime | Retrieve an artifact by digest into local storage |
| `validate` | plane → runtime | Re-verify digest, signature and compatibility locally |
| `load` | plane → runtime | Materialise a bundle inside an isolation boundary |
| `activate` | plane → runtime | Make skill exports invocable |
| `drain` | plane → runtime | Stop new invocations, wait for in-flight ones |
| `stop` | plane → runtime | Halt an instance after draining |
| `unload` | plane → runtime | Release the boundary and its resources |
| `quarantine` | plane → runtime | Isolate without discarding forensics |
| `report` | runtime → plane | Publish actual state |

Expose these as a **reconciliation** of desired state, not as an
imperative call sequence: the plane sends a generation plus a set of
bindings, and the runtime computes the plan.

## Instance lifecycle

```text
FETCHED -> VALIDATED -> LOADED -> ACTIVE -> DRAINING -> STOPPED -> UNLOADED
                                     |
                          FAILED / QUARANTINED
```

Only `ACTIVE` accepts new invocations. Every other state is fail-closed.

## Transport

SBRP is transport-neutral with JSON payloads. HTTP, gRPC, stdio, a
message bus or in-process calls are all valid. Transport is an adapter
concern and is invisible to the control-plane core.

## The six mandatory rules

1. **Double verification.** The runtime re-verifies artifact digest and
   signature at load time. Control-plane verification gates availability;
   it does not replace local verification. A runtime that trusts the
   plane's integrity claim is non-conformant.
2. **No runtime-side resolution.** The runtime executes a
   `ResolutionLock` and never selects versions itself, so composition
   stays replayable.
3. **Mediated invocation.** Even for a bundle in the caller's own
   process, the runtime mediates the call so authorisation, timeout,
   cancellation, tracing and resource accounting survive.
4. **Revocation is immediate.** Revoking stops new invocations at once,
   then drains. It does not wait for the next reconciliation pass.
5. **Fail-closed on partition.** Losing the control plane must never
   widen runtime authority. A partitioned runtime keeps serving what it
   was already authorised to serve and accepts nothing new.
6. **Honest reporting.** Report the state you are actually in, including
   `FAILED` and `QUARANTINED`. Reporting desired state as actual state is
   the most damaging conformance violation, because it silently defeats
   every drift and rollback mechanism above it.

## Implementing a runtime

```go
import syndovela "github.com/axisrobo/syndovela-open/sdk/go"

func describe() syndovela.RuntimeDescriptor {
    return syndovela.RuntimeDescriptor{
        RuntimeID:        "node-a7",
        Implementation:   "my-runtime",
        ProtocolVersions: []string{syndovela.ProtocolVersion},
        Isolation:        []string{"process"},
        ABIs:             []string{"native/grpc"},
        Platform:         "linux/amd64",
    }
}
```

Then implement the operation set and report `BundleInstance` records
honestly. Nothing else is required. Your runtime is eligible to host any
bundle whose declared protocol, ABI and isolation you advertise.

## Versioning

The version string is `sbrp/v<major>`. Additive optional fields are
allowed within a major version; anything breaking creates `sbrp/v2`, and
runtimes may advertise several versions at once during migration.

A breaking change requires review by at least one independently built
runtime implementation. A protocol only its author can implement is not
open.
