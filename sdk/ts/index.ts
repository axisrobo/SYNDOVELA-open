/**
 * SYNDOVELA TypeScript SDK
 *
 * Dependency-free client for the SYNDOVELA control plane, plus the SBRP
 * types runtime implementers need. Everything here mirrors the published
 * Go SDK; there is intentionally no runtime dependency.
 */

/** Bundle manifest types. */

export interface BundleMetadata {
  id: string;
  version: string;
  publisher: string;
  description?: string;
  labels?: Record<string, string>;
}

export interface Skill {
  id: string;
  contract: string;
  implementation: string;
  capabilityRefs?: string[];
  inputSchemaRef?: string;
  outputSchemaRef?: string;
  effectSchemaRef?: string;
}

export interface SkillRequirement {
  contract: string;
  version: string;
}

export interface BundleRequirement {
  id: string;
  version: string;
  reason?: string;
}

export interface Runtime {
  protocol: string;
  abi: string[];
  isolation?: string[];
  minIsolation?: string;
  platforms?: string[];
  features?: string[];
  constraints?: Record<string, string>;
  resources?: { cpu?: string; memory?: string };
}

export interface Security {
  signature: "required" | "optional";
  sbom: "required" | "optional";
  provenance: "required" | "optional";
  permissions?: string[];
}

export interface Artifact {
  digest: string;
  mediaType: string;
  size: number;
  platform?: string;
}

export interface Bundle {
  apiVersion: string;
  kind: "Bundle";
  metadata: BundleMetadata;
  skills: Skill[];
  requires?: { skills?: SkillRequirement[]; bundles?: BundleRequirement[] };
  runtime: Runtime;
  security: Security;
  artifacts?: Artifact[];
}

export interface ResolutionLock {
  lockId: string;
  resolverVersion: string;
  digest: string;
  inputDigest: string;
  selected: { bundleId: string; version: string; digest: string; reason?: string }[];
  bindings: { contract: string; skillId: string; bundleId: string; version: string }[];
}

export interface ConflictExplanation {
  ref: string;
  code: string;
  message: string;
}

/** SBRP runtime types. */

export interface RuntimeDescriptor {
  runtimeId: string;
  implementation?: string;
  implementationVersion?: string;
  protocolVersions: string[];
  isolation: string[];
  abis: string[];
  platform: string;
  features?: string[];
  limits?: Record<string, string>;
  labels?: Record<string, string>;
}

export const PROTOCOL_VERSION = "sbrp/v1";

export const INSTANCE_ACTIVE = "ACTIVE";
export type InstanceState =
  | "FETCHED" | "VALIDATED" | "LOADED" | "ACTIVE" | "DRAINING"
  | "STOPPED" | "UNLOADED" | "FAILED" | "QUARANTINED";

export interface BundleInstance {
  instanceId: string;
  runtimeId: string;
  nodeId?: string;
  bundleId: string;
  version: string;
  digest: string;
  state: InstanceState;
  isolation: string;
  activeInvocations: number;
  health?: string;
  loadedAt: string;
  reportedAt: string;
}

/** Control plane client. */

export class SyndovelaError extends Error {
  constructor(public statusCode: number, message: string) {
    super(message);
    this.name = "SyndovelaError";
  }
}

export class SyndovelaClient {
  private baseURL: string;

  constructor(baseURL: string, private fetchImpl: typeof fetch = fetch) {
    this.baseURL = baseURL.replace(/\/+$/, "");
  }

  async health(): Promise<{ product: string; version: string }> {
    return this.get("/healthz");
  }

  async registerBundle(bundle: Bundle): Promise<{ bundleId: string; version: string; digest: string; state: string }> {
    return this.post("/v1/bundles", bundle);
  }

  async resolve(requirements: unknown): Promise<{ resolutionLock: ResolutionLock; conflicts?: ConflictExplanation[] }> {
    return this.post("/v1/resolutions", requirements);
  }

  async applyRuntimeProfile(profile: unknown): Promise<{ name: string; generation: number }> {
    return this.post("/v1/runtime-profiles", profile);
  }

  private async request(method: string, path: string, body?: unknown): Promise<any> {
    const headers: Record<string, string> = { Accept: "application/json" };
    let payload: string | undefined;
    if (body !== undefined) {
      headers["Content-Type"] = "application/json";
      payload = JSON.stringify(body);
    }
    const resp = await this.fetchImpl(this.baseURL + path, { method, headers, body: payload });
    const text = await resp.text();
    if (resp.status < 200 || resp.status >= 300) {
      throw new SyndovelaError(resp.status, text);
    }
    return text ? JSON.parse(text) : undefined;
  }

  private get(path: string): Promise<any> {
    return this.request("GET", path);
  }

  private post(path: string, body: unknown): Promise<any> {
    return this.request("POST", path, body);
  }
}
