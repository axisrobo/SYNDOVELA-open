# Examples

| File | Shows |
| --- | --- |
| `identity-core-bundle.json` | A Bundle exporting three Skills, depending on another Skill by contract rather than by bundle id, declaring acceptable isolation levels and requiring full supply-chain metadata |
| `hr-offboarding-profile.json` | A Runtime Profile composing four Bundles with config, policy, placement and a default isolation level |

## Reading the Bundle example

- `skills[].contract` and `skills[].implementation` are different version
  axes. A contract may stay at `v2` across many implementation releases.
- `capabilityRefs` point at MODUREGIS capability versions. They
  reference; they never redefine.
- `requires.skills` names a contract, not a bundle. The resolver picks a
  qualifying provider, so the bundle is not hard-wired to one supplier.
- `runtime.isolation` lists what the bundle can tolerate. The runtime
  chooses and enforces the actual boundary.
- `security.permissions` must be complete. Undeclared access is denied.

## Reading the Runtime Profile example

- `generation` increments on every change. Deployments target a
  generation, not a floating profile.
- Version ranges here are resolved once into a ResolutionLock; the
  runtime never re-resolves them per invocation.
- `policyRefs` are references into AEGIVELA. SYNDOVELA stores the
  reference and never makes the authorisation decision itself.
