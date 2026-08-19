# Deployment

## Topologies

| Domain | Components | Principle |
| --- | --- | --- |
| Developer / local | CLI, local registry and cache, local resolver | Dev fallbacks are allowed, but any signature or permission bypass must be explicitly dev-only |
| Enterprise control plane | Registry, resolver, profile manager, deployment controller, supply chain, console | HA, multi-tenant, audited, with federated mirrors |
| Datacentre runtime | Regional cache and mirror | Losing the control plane must never widen runtime authority |
| Edge / robot | Offline resolution pack and signed bundle cache | Locks and bundle versions must verify without network access |
| Federated ecosystem | Cross-organisation metadata federation | No shared database; interoperate through signed manifests, digests and policy references |

## Requirements

- PostgreSQL 18 or later for control-plane state.
- Persistent storage for artifacts, or an OCI-compatible registry.
- Outbound access to any configured mirrors and trust roots.

## Configuration

| Variable | Purpose | Default |
| --- | --- | --- |
| `SYNDOVELA_LISTEN_ADDR` | HTTP listen address | `:8080` |
| `DATABASE_URL` | PostgreSQL connection string | required |

## Migrations

Migrations are forward-only and strictly additive, so they are safe to
apply online. Run `syndovela-migrate` before starting a new binary
version.

## Scale targets

Validated before GA:

| Dimension | Target |
| --- | --- |
| Skill metadata records | 100,000 |
| Bundles | 10,000 |
| Runtime Profiles | 1,000 |
| Runtime nodes | 10,000 |
| Bundle instances | 100,000 |
| Resolution p95 (cache hit) | under 500 ms |

## Operational rules

1. Deployments reference a `ResolutionLock`, never a floating range.
2. Upgrades activate the new version before draining the old one, so
   in-flight invocations are never orphaned.
3. Rollback points at a previous lock. Published bundle versions are
   immutable and are never rewritten.
4. Revocation stops new invocations immediately, then drains. It does not
   wait for the next rollout window.
5. Control-plane outage must degrade to "no new compositions", never to
   "unverified compositions".
