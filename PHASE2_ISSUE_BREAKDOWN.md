# Phase 2 Issue Breakdown

This document breaks the Phase 2 epics from `ISSUE_BACKLOG.md` into smaller implementable issues. Child issue IDs stay directly linked to the parent epic and mirror the GitHub issue structure.

## Convention

- Parent epic: `ISS-101-S`
- Child issue: `ISS-101-S-01`, `ISS-101-S-02`, `ISS-101-S-03`

## Delivery Focus

Phase 2 builds on the stable CLI-first core from Phase 1. The main goals are:

1. Introduce a practical TUI without duplicating business logic.
2. Add MCP and extension mechanisms for tools and commands.
3. Introduce declarative agents, skills, and orchestration.
4. Add controlled remote access, routing, and model update workflows.

## Recommended Delivery Waves

### Wave 1 - Interface and extension foundations

- `ISS-101-S-01` Create TUI app shell and navigation state
- `ISS-102-S-01` Define recommended MCP and toolchain profiles with dependency probes
- `ISS-103-S-01` Define MCP runtime contracts, config model, and session lifecycle
- `ISS-104-S-01` Define extension manifest, loader, and policy model
- `ISS-105-S-01` Define declarative schema for skills and agent definitions
- `ISS-108-S-01` Define transport-neutral API and SDK contract with versioning policy
- `ISS-109-S-01` Define model version metadata, update policy, and compatibility rules

### Wave 2 - Core capabilities on top of the Phase 1 runtime

- `ISS-101-S-02` Wire TUI chat view to existing run and chat services
- `ISS-101-S-03` Add TUI panels for model switch, history, and resource status
- `ISS-103-S-02` Implement MCP client adapter and tool or resource bridge
- `ISS-104-S-02` Load custom commands into CLI and TUI command surfaces
- `ISS-104-S-03` Load custom tools into the runtime with policy gating
- `ISS-105-S-02` Implement filesystem loader and validator for agent assets
- `ISS-105-S-03` Bind loaded skills and declarative agents into execution runtime
- `ISS-107-S-01` Define structured tool schemas, validation, and tool error model
- `ISS-108-S-02` Implement authenticated server mode on top of application services
- `ISS-109-S-02` Implement update check and surfaced status in CLI, TUI, and SDK

### Wave 3 - Orchestration and automation

- `ISS-102-S-02` Build guided setup flow for MCP and standard tools
- `ISS-103-S-03` Expose selected Noctua capabilities through an MCP server
- `ISS-106-S-01` Define role contracts, handoff envelope, and orchestration trace model
- `ISS-106-S-02` Implement the four-role orchestration pipeline
- `ISS-106-S-03` Add policy-based routing for local, remote, and tool-backed execution
- `ISS-107-S-02` Implement local-model tool-calling loop with guarded execution
- `ISS-107-S-03` Implement controlled external worker launcher and result handoff
- `ISS-108-S-03` Implement remote client CLI and thin SDK wrapper
- `ISS-109-S-03` Implement atomic model update workflow with rollback

### Wave 4 - Hardening and test coverage

- `ISS-101-S-04` Harden TUI UX with loading, error states, and navigation tests
- `ISS-102-S-03` Add post-setup validation and repair command
- `ISS-103-S-04` Add MCP integration tests, audit logs, and policy hooks
- `ISS-104-S-04` Add extension diagnostics, audit events, and regression tests
- `ISS-105-S-04` Add example skills, agents, and contract tests
- `ISS-106-S-04` Add orchestration replay, evaluation reports, and end-to-end tests
- `ISS-107-S-04` Add timeout, cleanup, and audit tests for tools and workers
- `ISS-108-S-04` Add compatibility tests, migration handling, and security hardening
- `ISS-109-S-04` Preserve alias mappings and add update audit and regression tests

## Detailed Breakdown

### ISS-101-S TUI shell for chat, history, and models

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-101-S-01` | Create TUI app shell and navigation state | P1 | `ISS-002-M-02`, `ISS-003-M-03`, `ISS-010-M-03` | Stand up the base TUI frame on top of the existing Phase 1 services. |
| `ISS-101-S-02` | Wire TUI chat view to existing run and chat services | P1 | `ISS-101-S-01`, `ISS-002-M-02`, `ISS-007-M-03`, `ISS-006-M-03` | Make chat usable inside the TUI. |
| `ISS-101-S-03` | Add TUI panels for model switch, history, and resource status | P1 | `ISS-101-S-01`, `ISS-004-M-03`, `ISS-006-M-02`, `ISS-007-M-03`, `ISS-010-M-03` | Expose key management flows inside the TUI. |
| `ISS-101-S-04` | Harden TUI UX with loading, error states, and navigation tests | P2 | `ISS-101-S-02`, `ISS-101-S-03` | Make the TUI stable enough for daily use. |

### ISS-102-S Guided setup for MCP and standard tools

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-102-S-01` | Define recommended MCP and toolchain profiles with dependency probes | P1 | `ISS-003-M-03` | Standardize recommended setup profiles. |
| `ISS-102-S-02` | Build guided setup flow for MCP and standard tools | P1 | `ISS-102-S-01`, `ISS-103-S-02`, `ISS-104-S-01` | Let users enable the standard toolchain without manual setup. |
| `ISS-102-S-03` | Add post-setup validation and repair command | P2 | `ISS-102-S-02`, `ISS-103-S-04`, `ISS-104-S-04` | Make onboarding supportable and recoverable. |

### ISS-103-S MCP client and server integration

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-103-S-01` | Define MCP runtime contracts, config model, and session lifecycle | P1 | `ISS-001-M-01`, `ISS-003-M-03` | Add MCP as a first-class runtime boundary. |
| `ISS-103-S-02` | Implement MCP client adapter and tool or resource bridge | P1 | `ISS-103-S-01` | Let Noctua consume external MCP servers. |
| `ISS-103-S-03` | Expose selected Noctua capabilities through an MCP server | P1 | `ISS-103-S-01` | Let external MCP clients call into Noctua. |
| `ISS-103-S-04` | Add MCP integration tests, audit logs, and policy hooks | P2 | `ISS-103-S-02`, `ISS-103-S-03` | Make MCP support traceable and safe. |

### ISS-104-S Registry for custom commands and tools

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-104-S-01` | Define extension manifest, loader, and policy model | P1 | `ISS-001-M-01`, `ISS-003-M-03` | Create a shared registration model for custom commands and tools. |
| `ISS-104-S-02` | Load custom commands into CLI and TUI command surfaces | P1 | `ISS-104-S-01`, `ISS-101-S-01` | Let users add commands without core changes. |
| `ISS-104-S-03` | Load custom tools into the runtime with policy gating | P1 | `ISS-104-S-01` | Make custom tools available to tool-calling and agent flows. |
| `ISS-104-S-04` | Add extension diagnostics, audit events, and regression tests | P2 | `ISS-104-S-02`, `ISS-104-S-03` | Make the extension system maintainable. |

### ISS-105-S Agent skills and declarative agent definitions

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-105-S-01` | Define declarative schema for skills and agent definitions | P1 | `ISS-001-M-01`, `ISS-003-M-03` | Standardize how skills and agents are authored. |
| `ISS-105-S-02` | Implement filesystem loader and validator for agent assets | P1 | `ISS-105-S-01`, `ISS-104-S-01` | Load agents and skills without code changes. |
| `ISS-105-S-03` | Bind loaded skills and declarative agents into execution runtime | P1 | `ISS-105-S-02` | Make declarative agents executable. |
| `ISS-105-S-04` | Add example skills, agents, and contract tests | P2 | `ISS-105-S-03` | Lock in authoring and runtime behavior. |

### ISS-106-S Multi-agent orchestration and routing

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-106-S-01` | Define role contracts, handoff envelope, and orchestration trace model | P1 | `ISS-105-S-01` | Create a stable execution model for multi-agent flows. |
| `ISS-106-S-02` | Implement the four-role orchestration pipeline | P1 | `ISS-106-S-01`, `ISS-105-S-03` | Run tasks through the full multi-agent workflow. |
| `ISS-106-S-03` | Add policy-based routing for local, remote, and tool-backed execution | P1 | `ISS-106-S-01`, `ISS-103-S-02`, `ISS-107-S-02` | Make orchestration choose the right path per task. |
| `ISS-106-S-04` | Add orchestration replay, evaluation reports, and end-to-end tests | P2 | `ISS-106-S-02`, `ISS-106-S-03` | Make multi-agent behavior debuggable. |

### ISS-107-S Better tool calling for local models and worker launch

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-107-S-01` | Define structured tool schemas, validation, and tool error model | P1 | `ISS-001-M-01`, `ISS-104-S-01` | Make tool calling reliable for local models. |
| `ISS-107-S-02` | Implement local-model tool-calling loop with guarded execution | P1 | `ISS-107-S-01`, `ISS-104-S-03` | Let local models call tools safely. |
| `ISS-107-S-03` | Implement controlled external worker launcher and result handoff | P1 | `ISS-107-S-01`, `ISS-104-S-03` | Support safe startup of external AI workers from tools. |
| `ISS-107-S-04` | Add timeout, cleanup, and audit tests for tools and workers | P2 | `ISS-107-S-02`, `ISS-107-S-03` | Prevent tool execution from leaving bad state behind. |

### ISS-108-S Server mode, remote client, and SDK or API

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-108-S-01` | Define transport-neutral API and SDK contract with versioning policy | P1 | `ISS-001-M-01`, `ISS-002-M-02` | Expose core capabilities without coupling callers to CLI internals. |
| `ISS-108-S-02` | Implement authenticated server mode on top of application services | P1 | `ISS-108-S-01`, `ISS-003-M-03`, `ISS-010-M-03` | Run Noctua as a secure local or remote service. |
| `ISS-108-S-03` | Implement remote client CLI and thin SDK wrapper | P1 | `ISS-108-S-01`, `ISS-108-S-02` | Make server mode usable from other clients and apps. |
| `ISS-108-S-04` | Add compatibility tests, migration handling, and security hardening | P2 | `ISS-108-S-02`, `ISS-108-S-03` | Keep the remote interface stable across upgrades. |

### ISS-109-S Model updates and version management

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-109-S-01` | Define model version metadata, update policy, and compatibility rules | P1 | `ISS-004-M-01`, `ISS-005-M-02` | Extend model lifecycle management with update awareness. |
| `ISS-109-S-02` | Implement update check and surfaced status in CLI, TUI, and SDK | P1 | `ISS-109-S-01`, `ISS-004-M-02`, `ISS-101-S-03`, `ISS-108-S-03` | Let users detect available updates. |
| `ISS-109-S-03` | Implement atomic model update workflow with rollback | P1 | `ISS-109-S-01`, `ISS-109-S-02`, `ISS-005-M-02` | Apply updates safely without breaking installs. |
| `ISS-109-S-04` | Preserve alias mappings and add update audit and regression tests | P2 | `ISS-109-S-03` | Protect user-facing behavior across upgrades. |

## Critical Path

The most important Phase 2 dependency chain is:

1. `ISS-101-S-01`
2. `ISS-103-S-01`
3. `ISS-104-S-01`
4. `ISS-105-S-01`
5. `ISS-105-S-02`
6. `ISS-105-S-03`
7. `ISS-106-S-01`
8. `ISS-107-S-01`
9. `ISS-107-S-02`
10. `ISS-106-S-02`
11. `ISS-106-S-03`
12. `ISS-108-S-01`
13. `ISS-108-S-02`
14. `ISS-108-S-03`

The TUI and onboarding tracks can advance in parallel after the shared extension and MCP foundations exist.
