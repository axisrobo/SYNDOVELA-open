# Go SDK

```powershell
go get github.com/axisrobo/syndovela-open/sdk/go
```

The SDK has zero third-party dependencies and is licensed Apache-2.0.

## Client

```go
c := syndovela.New("http://localhost:8080")

// custom transport, timeouts, auth round-tripper
c = syndovela.New(baseURL, syndovela.WithHTTPClient(&http.Client{
    Timeout: 10 * time.Second,
}))
```

| Method | Endpoint |
| --- | --- |
| `Health(ctx)` | `GET /healthz` |
| `RegisterBundle(ctx, bundle)` | `POST /v1/bundles` |
| `Resolve(ctx, requirements)` | `POST /v1/resolutions` |
| `ApplyRuntimeProfile(ctx, profile)` | `POST /v1/runtime-profiles` |

## Errors

Non-2xx responses return `*APIError` carrying the status code and the raw
body, so callers can distinguish an immutability conflict (409) from an
unresolvable requirement set (422).

```go
if apiErr, ok := err.(*syndovela.APIError); ok {
    switch apiErr.StatusCode {
    case http.StatusConflict:
        // that bundle version is already published and immutable
    case http.StatusUnprocessableEntity:
        // read the conflict explanations in apiErr.Body
    }
}
```

## SBRP types

The SDK also publishes the Skill Bundle Runtime Protocol types
(`RuntimeDescriptor`, `BundleBinding`, `BundleInstance`,
`SkillInvocation`, `ActualStateReport`) so that runtime authors can
implement the protocol without depending on any control-plane
implementation or on AGPL source.

```go
func (r *myRuntime) Describe() syndovela.RuntimeDescriptor {
    return syndovela.RuntimeDescriptor{
        RuntimeID:        r.id,
        ProtocolVersions: []string{syndovela.ProtocolVersion},
        Isolation:        []string{"process"},
        ABIs:             []string{"native/grpc"},
        Platform:         "linux/amd64",
    }
}
```

Advertise only what you can genuinely enforce. The control plane places
bundles based on these claims, so over-advertising isolation is a
security defect rather than an optimism bug.

## Types

`Bundle`, `Metadata`, `Skill`, `Requires`, `Runtime`, `Security`,
`Artifact`, `RuntimeProfile` and `ResolutionLock` mirror the published
contract. Keep the three version axes distinct: `Skill.Contract`,
`Skill.Implementation` and `Metadata.Version` are independent, and none
implies another.

## Context

Every call takes a `context.Context`. Resolution over a large dependency
graph can be slow; always pass a deadline.
