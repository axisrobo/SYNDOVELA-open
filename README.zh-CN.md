# SYNDOVELA Open

[English](README.md)

> SYNDOVELA 的开放 SDK、协议、示例和已签名二进制发行物。SYNDOVELA 是用于跨运行时治理 Agent Skill 组合的企业平台。

## SYNDOVELA

SYNDOVELA 是一个在企业规模下编排和治理可复用 Agent Skill 的控制平面。它位于
Skill 资产与执行这些 Skill 的运行时之间：将 Skill 打包为可签名、可版本化、可组合、
可解析、可分发和可撤销的 Bundle 与 Runtime Profile。运行时仍负责加载、隔离、激活、
调用、观测和卸载。

### 要解决的问题

在企业规模下，将每个可复用 Skill 作为独立服务部署，会把代码复用问题变成运维问题：

```text
10,000 个 Skill -> 10,000 个服务 -> 10,000 个容器 -> 10,000 个运维边界
```

SYNDOVELA 在不减少 Skill 数量的前提下，缩小治理面：

```text
10,000 个 Skill -> 约 1,000 个 Bundle -> 约 100 个 Runtime Profile -> 共享运行时基座
```

### 能力与价值

- **部署蔓延**：相关 Skill 共享一个 Bundle 生命周期，而不必各自维护部署边界。
- **版本漂移**：`ResolutionLock` 固定运行时应持有的确切 Bundle 版本和摘要；调用时不会重新解析。
- **供应链不透明**：每个 Bundle 可携带签名、SBOM 和来源证明；存在于注册表中不等于已获批准。
- **不安全升级**：新版本在旧版本排空前激活；回滚只需指向先前的 lock，而非修改发布物。
- **无治理的撤销**：撤销 Bundle 或发布者可立即阻止整个集群中的新调用。

### 开放设计

SYNDOVELA 通过 **SBRP**（Skill Bundle Runtime Protocol）与运行时通信。SBRP 是一项开放、
厂商中立的协议。

- Bundle 声明 `protocol`、`abi` 和 `isolation`，而非厂商；同一个 Bundle 可运行于任何符合规范的运行时。
- 兼容性取决于已声明的能力，而不是运行时身份。
- 控制平面与运行时均可独立实现和测试。

参见 [`docs/sbrp.md`](docs/sbrp.md) 和 [`docs/conformance.md`](docs/conformance.md)。

## SYNDOVELA-open

本仓库以 Apache-2.0 许可证发布，是 SYNDOVELA 的公共发行与集成入口。它包含开放协议和
API 契约、SDK、示例、一致性材料以及已签名的发行二进制文件。仓库不包含 SYNDOVELA 核心源
代码；采用 AGPL-3.0-or-later 许可证的控制平面实现位于
[SYNDOVELA core repository](https://github.com/axisrobo/SYNDOVELA)。

| 路径 | 内容 |
| --- | --- |
| `docs/sbrp.md` | Skill Bundle Runtime Protocol |
| `docs/conformance.md` | 控制平面和运行时的一致性检查清单 |
| `docs/cli.md` | `syndovela-cli` 参考（`pack`、`pack-verify`、`warm`、`impact`、`bench`、`conform`） |
| `api/` | 控制平面 API 的已发布 OpenAPI 描述，与 core v1.9.0 保持一致 |
| `sdk/go/` | 无依赖的 Go 客户端和 SBRP 类型 |
| `examples/` | Bundle 清单、Runtime Profile 和运行时描述符 |
| `console/` | 消费 `/v1/stats` 的只读 EE 仪表板脚手架 |
| `releases/` | 已签名的 core 二进制文件、校验和与 SBOM；每个发行版均记录所含 core 版本 |
| `scripts/` | `verify-release` 校验和与签名验证工具 |

## 核心概念

| 对象 | 含义 |
| --- | --- |
| `Skill` | 调用单元：一个可调用操作的契约 |
| `Bundle` | 一组相关 Skill 的打包、版本与治理单元 |
| `BundleVersion` | Bundle 的不可变、按摘要寻址的发行版本 |
| `ResolutionLock` | 一次解析的确定性依赖闭包 |
| `RuntimeProfile` | 一个运行时域的期望组合 |
| `BundleInstance` | 由运行时生成并通过 SBRP 上报的实际状态 |

## 兼容性

SYNDOVELA 使用 `major.minor.patch` 版本号。Bundle 契约和 SBRP 与产品独立版本化：破坏性
变更会进入新的契约主版本，旧版本会被冻结而不会被修改。运行时可以在迁移期间同时声明多个
SBRP 版本。

## 许可证

[Apache-2.0](LICENSE)。随仓库分发的 core 二进制文件附带其自身的 AGPL-3.0-or-later 声明。
