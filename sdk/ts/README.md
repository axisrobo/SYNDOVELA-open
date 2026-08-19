# TypeScript SDK

Dependency-free TypeScript client for the SYNDOVELA control plane and
the SBRP runtime types.

```ts
import { SyndovelaClient, PROTOCOL_VERSION } from "@axisrobo/syndovela";

const client = new SyndovelaClient("http://localhost:8080");

await client.registerBundle({
  apiVersion: "syndovela.axisrobo.io/v1",
  kind: "Bundle",
  metadata: { id: "example.identity-core", version: "0.1.0", publisher: "example" },
  skills: [{ id: "identity.lookup", contract: "identity.lookup/v2", implementation: "0.1.0" }],
  runtime: { protocol: PROTOCOL_VERSION, abi: ["wasi/preview2"], isolation: ["wasm"] },
  security: { signature: "required", sbom: "required", provenance: "required" },
});

try {
  await client.resolve({ skills: [{ ref: "identity.lookup/v2" }] });
} catch (err) {
  if (err instanceof SyndovelaError && err.statusCode === 422) {
    // registry presence is not approval: a REGISTERED bundle is not deployable
  }
}
```

Errors from the control plane are `SyndovelaError` with the HTTP status
code, so 409 (immutable version), 422 (validation/resolution) and 404
can be distinguished.

## SBRP runtime types

`RuntimeDescriptor`, `BundleInstance`, `InstanceState` and
`PROTOCOL_VERSION` let runtime authors implement the protocol with no
dependency on a control plane implementation.
