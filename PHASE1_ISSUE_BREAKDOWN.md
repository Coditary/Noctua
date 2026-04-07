# Phase 1 Issue Breakdown

This document breaks the Phase 1 epics from `ISSUE_BACKLOG.md` into smaller implementable issues. Child issue IDs stay directly linked to the parent epic.

## Convention

- Parent epic: `ISS-001-M`
- Child issue: `ISS-001-M-01`, `ISS-001-M-02`, `ISS-001-M-03`

## Recommended Delivery Waves

### Wave 1 - Core foundation

- `ISS-001-M-01` Define core runtime contracts and adapter interfaces
- `ISS-003-M-01` Define config schema, profile model and precedence rules
- `ISS-003-M-02` Implement local path resolver and app directory bootstrap
- `ISS-010-M-01` Define structured log and event schema plus redaction rules

### Wave 2 - Provider and model foundations

- `ISS-001-M-02` Implement provider registry and alias resolution foundation
- `ISS-001-M-03` Add compliance tests and starter adapters for one local and one remote provider
- `ISS-004-M-01` Define model catalog and installed model metadata
- `ISS-004-M-02` Implement available-model listing and search across configured sources
- `ISS-003-M-03` Implement config loading and local storage abstractions

### Wave 3 - Lifecycle and runtime

- `ISS-005-M-01` Wrap the existing download tool behind an install adapter
- `ISS-005-M-02` Implement install state machine with staging and atomic promotion
- `ISS-006-M-01` Implement model selection resolution for global, profile and call scopes
- `ISS-008-M-01` Define AI parameter schema and validation rules
- `ISS-008-M-02` Implement parameter merge logic for config, profile and CLI overrides
- `ISS-007-M-01` Define session and message storage schema
- `ISS-007-M-02` Implement chat session persistence and resume service

### Wave 4 - User-facing CLI flows

- `ISS-002-M-01` Define CLI command surface, global flags and output modes
- `ISS-002-M-02` Implement `run` and `chat` commands against application services
- `ISS-002-M-03` Implement `status` command with exit codes and JSON output
- `ISS-004-M-03` Implement installed-model listing and merged catalog presentation
- `ISS-005-M-03` Add CLI install flow and install job reporting
- `ISS-006-M-02` Persist active model mappings and alias targets
- `ISS-006-M-03` Expose runtime context output for model, provider, profile and thinking state
- `ISS-007-M-03` Add CLI support for resume, history listing and continue flow
- `ISS-008-M-03` Expose parameter override flags and resolved parameter inspection
- `ISS-009-M-03` Add CLI delete command with non-interactive safety flags

### Wave 5 - Cleanup, observability, and hardening

- `ISS-009-M-01` Define uninstall safety checks and artifact ownership rules
- `ISS-009-M-02` Implement delete workflow with cleanup and failure recovery
- `ISS-010-M-02` Implement resource probes for CPU, RAM and GPU status
- `ISS-010-M-03` Wire logs, metrics, and resource snapshots into CLI and model lifecycle flows

## Detailed Breakdown

### ISS-001-M Core runtime and provider abstraction

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-001-M-01` | Define core runtime contracts and adapter interfaces | P1 | - | Define shared ports and core models for local and external execution. |
| `ISS-001-M-02` | Implement provider registry and alias resolution foundation | P1 | `ISS-001-M-01` | Make providers and alias targets centrally registrable and resolvable. |
| `ISS-001-M-03` | Add compliance tests and starter adapters for one local and one remote provider | P1 | `ISS-001-M-01`, `ISS-001-M-02` | Prove the core works with at least one local and one external adapter. |

### ISS-002-M CLI foundation for run, chat, and status

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-002-M-01` | Define CLI command surface, global flags and output modes | P1 | `ISS-001-M-01`, `ISS-003-M-01` | Define a stable command surface for the MVP. |
| `ISS-002-M-02` | Implement `run` and `chat` commands against application services | P1 | `ISS-002-M-01`, `ISS-001-M-03`, `ISS-006-M-01`, `ISS-008-M-02` | Build the first production user flows on top of the core. |
| `ISS-002-M-03` | Implement `status` command with exit codes and JSON output | P2 | `ISS-002-M-01`, `ISS-003-M-03` | Provide machine-readable status output and clean exit codes. |

### ISS-003-M Local configuration, profiles, and storage layout

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-003-M-01` | Define config schema, profile model and precedence rules | P1 | - | Define effective defaults, profiles, and override rules. |
| `ISS-003-M-02` | Implement local path resolver and app directory bootstrap | P1 | `ISS-003-M-01` | Add configurable local directory layout and bootstrap behavior. |
| `ISS-003-M-03` | Implement config loading and local storage abstractions | P1 | `ISS-003-M-01`, `ISS-003-M-02` | Create reusable access to config, data, cache, and logs. |

### ISS-004-M Model catalog for installed and available models

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-004-M-01` | Define model catalog, catalog entry and installed model metadata | P1 | `ISS-001-M-01` | Define a shared data model for installed and available models. |
| `ISS-004-M-02` | Implement available-model listing and search across configured sources | P1 | `ISS-004-M-01`, `ISS-001-M-03` | Allow catalogs from multiple sources to be searched and listed. |
| `ISS-004-M-03` | Implement installed-model listing and merged catalog presentation | P2 | `ISS-004-M-01`, `ISS-003-M-03` | Show installed and available models consistently. |

### ISS-005-M Automated model installation and download pipeline

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-005-M-01` | Wrap the existing download tool behind an install adapter | P1 | `ISS-001-M-01`, `ISS-004-M-01` | Plug existing download logic into the new architecture. |
| `ISS-005-M-02` | Implement install state machine with staging and atomic promotion | P1 | `ISS-005-M-01`, `ISS-003-M-03` | Ensure consistent installation without broken intermediate states. |
| `ISS-005-M-03` | Add CLI install flow and install job reporting | P2 | `ISS-002-M-01`, `ISS-005-M-02` | Expose model installation to end users through the CLI. |

### ISS-006-M Model selection and transparent runtime context

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-006-M-01` | Implement model selection resolution for global, profile and call scopes | P1 | `ISS-001-M-02`, `ISS-003-M-03`, `ISS-004-M-01` | Define how the active model is resolved for each scope. |
| `ISS-006-M-02` | Persist active model mappings and alias targets | P2 | `ISS-006-M-01`, `ISS-003-M-03` | Make active selections durable across sessions. |
| `ISS-006-M-03` | Expose runtime context output for model, provider, profile and thinking state | P2 | `ISS-006-M-01`, `ISS-008-M-02` | Show the effective runtime context transparently to users. |

### ISS-007-M Persist chat history and resume sessions

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-007-M-01` | Define session and message storage schema | P1 | `ISS-003-M-03` | Define the local model for sessions, messages, and turns. |
| `ISS-007-M-02` | Implement chat session persistence and resume service | P1 | `ISS-007-M-01`, `ISS-003-M-03` | Allow sessions to be stored, loaded, and resumed. |
| `ISS-007-M-03` | Add CLI support for resume, history listing and continue flow | P2 | `ISS-002-M-02`, `ISS-007-M-02` | Make session continuation and history access available to users. |

### ISS-008-M Parameter management and CLI overrides

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-008-M-01` | Define AI parameter schema and validation rules | P1 | `ISS-003-M-01` | Define one consistent parameter model for model calls. |
| `ISS-008-M-02` | Implement parameter merge logic for config, profile and CLI overrides | P1 | `ISS-008-M-01`, `ISS-003-M-03` | Compute effective parameters deterministically from multiple sources. |
| `ISS-008-M-03` | Expose parameter override flags and resolved parameter inspection | P2 | `ISS-002-M-01`, `ISS-008-M-02` | Make parameters overridable and inspectable at runtime. |

### ISS-009-M Model deletion and clean cleanup

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-009-M-01` | Define uninstall safety checks and artifact ownership rules | P1 | `ISS-004-M-01`, `ISS-005-M-02` | Define when a model can be removed safely. |
| `ISS-009-M-02` | Implement delete workflow with cleanup and failure recovery | P1 | `ISS-009-M-01`, `ISS-003-M-03` | Implement model deletion with clean cleanup and recovery. |
| `ISS-009-M-03` | Add CLI delete command with non-interactive safety flags | P2 | `ISS-002-M-01`, `ISS-009-M-02` | Expose deletion to users and scripts. |

### ISS-010-M Resource visibility, logging, and baseline observability

| Child ID | Title | Priority | Depends on | Goal |
| --- | --- | --- | --- | --- |
| `ISS-010-M-01` | Define structured log and event schema plus redaction rules | P1 | `ISS-001-M-01`, `ISS-003-M-01` | Structure logs and events cleanly from the beginning. |
| `ISS-010-M-02` | Implement resource probes for CPU, RAM and GPU status | P2 | `ISS-002-M-03`, `ISS-003-M-02` | Capture runtime resources for status and diagnostics. |
| `ISS-010-M-03` | Wire logs, metrics, and resource snapshots into CLI and model lifecycle flows | P2 | `ISS-010-M-01`, `ISS-010-M-02`, `ISS-002-M-02`, `ISS-005-M-03`, `ISS-009-M-03` | Integrate observability into real user and model lifecycle flows. |

## Critical Path

The tightest Phase 1 chain is:

1. `ISS-001-M-01`
2. `ISS-003-M-01`
3. `ISS-003-M-02`
4. `ISS-003-M-03`
5. `ISS-001-M-02`
6. `ISS-001-M-03`
7. `ISS-004-M-01`
8. `ISS-006-M-01`
9. `ISS-008-M-01`
10. `ISS-008-M-02`
11. `ISS-002-M-01`
12. `ISS-002-M-02`

After that, installation, history, deletion, and observability can be parallelized.
