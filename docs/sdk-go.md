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

## Types

`Bundle`, `Metadata`, `Skill`, `Requires`, `Runtime`, `Security`,
`Artifact`, `RuntimeProfile` and `ResolutionLock` mirror the published
contract. Keep the three version axes distinct: `Skill.Contract`,
`Skill.Implementation` and `Metadata.Version` are independent, and none
implies another.

## Context

Every call takes a `context.Context`. Resolution over a large dependency
graph can be slow; always pass a deadline.
