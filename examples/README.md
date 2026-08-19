# Examples

| File | Shows |
| --- | --- |
| `identity-core-bundle.json` | A Bundle exporting three Skills, depending on another Skill by contract rather than by bundle id, declaring vendor-neutral runtime requirements and requiring full supply-chain metadata |
| `hr-offboarding-profile.json` | A Runtime Profile composing four Bundles with config, policy, placement and a default isolation level |
| `runtime-descriptor.json` | What a runtime returns from SBRP `describe`, and the only information a control plane may use to decide eligibility |

## Reading the Bundle example

- `skills[].contract` and `skills[].implementation` are different version
  axes. A contract may stay at `v2` across many implementation releases.
- `capabilityRefs` point at an external capability registry. They
  reference; they never redefine.
- `requires.skills` names a contract, not a bundle. The resolver picks a
  qualifying provider, so the bundle is not hard-wired to one supplier.
- `runtime.protocol` and `runtime.abi` state what the bundle needs
  without naming a runtime product. Any runtime advertising a matching
  protocol version and ABI may host it.
- `runtime.isolation` lists what the bundle can tolerate, in preference
  order. The runtime chooses and enforces the actual boundary.
- `security.permissions` must be complete. Undeclared access is denied.

## Reading the runtime descriptor

This is the whole basis for compatibility. `implementation` and
`implementationVersion` are present for operators and logs only; a
control plane that branches on them has broken the protocol's openness
guarantee, because it makes unknown runtimes ineligible by default.

## Reading the Runtime Profile example

- `generation` increments on every change. Deployments target a
  generation, not a floating profile.
- Version ranges here are resolved once into a ResolutionLock; the
  runtime never re-resolves them per invocation.
- `policyRefs` are opaque references into whatever policy decision point
  you operate. SYNDOVELA stores the reference and never makes the
  authorisation decision itself.
