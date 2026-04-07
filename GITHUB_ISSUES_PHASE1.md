# GitHub Issues Phase 1

This document contains copy-paste-ready GitHub issue drafts for the Phase 1 child issues from `PHASE1_ISSUE_BREAKDOWN.md`.

## Template

Recommended labels:

- `phase:1`
- `priority:p1` or `priority:p2`
- `area:*`
- `type:*`

---

## ISS-001-M-01 Define core runtime contracts and adapter interfaces

Labels: `phase:1`, `priority:p1`, `area:core`, `type:architecture`

Parent: `ISS-001-M`

Related requirements: `F-001-M`, `F-004-M`, `F-015-M`, `NF-009-M`, `NF-010-M`

### Goal

Define the minimal stable runtime contracts and adapter interfaces for local engines and remote providers.

### Description

Noctua needs one shared execution model before building CLI flows and model lifecycle logic. This issue defines the core contracts that isolate the application layer from concrete providers and engines.

### Done

- Core request and result contracts are defined.
- Provider and engine adapter interfaces are defined.
- Model alias and capability concepts are represented in the core model.
- No provider-specific fields leak into the core contracts.

### Notes

Keep the first version minimal and focused on text and chat only.

---

## ISS-001-M-02 Implement provider registry and alias resolution foundation

Labels: `phase:1`, `priority:p1`, `area:core`, `type:feature`

Parent: `ISS-001-M`

Depends on: `ISS-001-M-01`

Related requirements: `F-004-M`, `F-009-M`, `F-015-M`

### Goal

Allow Noctua to register providers and resolve user-facing model aliases to concrete targets.

### Description

Users should be able to refer to a stable alias while the runtime maps it to a provider, engine, and concrete model target. This is the foundation for defaults, profiles, and later routing.

### Done

- A provider registry exists.
- A model alias can resolve to a concrete provider or engine target.
- Alias resolution failures return clear typed errors.
- The design supports global, profile, and per-call overrides.

---

## ISS-001-M-03 Add compliance tests and starter adapters for one local and one remote provider

Labels: `phase:1`, `priority:p1`, `area:providers`, `type:test`

Parent: `ISS-001-M`

Depends on: `ISS-001-M-01`, `ISS-001-M-02`

Related requirements: `F-001-M`, `F-015-M`, `NF-010-M`

### Goal

Prove that the adapter model works with at least one local and one remote provider.

### Description

The architecture is only useful if the same internal request path can reach both a local and an external provider. Add starter adapters and a shared compliance test suite.

### Done

- One local adapter exists.
- One remote adapter exists.
- Both pass the same compliance expectations for basic text and chat flows.
- Adapter extension points are documented well enough for follow-up provider work.

---

## ISS-002-M-01 Define CLI command surface, global flags and output modes

Labels: `phase:1`, `priority:p1`, `area:cli`, `type:architecture`

Parent: `ISS-002-M`

Depends on: `ISS-001-M-01`, `ISS-003-M-01`

Related requirements: `F-002-M`, `F-003-M`, `F-016-M`, `NF-008-M`

### Goal

Define the stable command surface for the Phase 1 CLI.

### Description

Before implementing commands, the CLI needs a clear shape for commands, global flags, profile selection, output format, and exit-code behavior.

### Done

- Core commands for `run`, `chat`, `models`, `status`, and config or profile selection are defined.
- Global flags and output mode behavior are documented.
- JSON or another machine-readable output mode is included where relevant.
- Exit-code semantics are defined for success, validation, and runtime failures.

---

## ISS-002-M-02 Implement `run` and `chat` commands against application services

Labels: `phase:1`, `priority:p1`, `area:cli`, `type:feature`

Parent: `ISS-002-M`

Depends on: `ISS-002-M-01`, `ISS-001-M-03`, `ISS-006-M-01`, `ISS-008-M-02`

Related requirements: `F-002-M`, `F-003-M`, `NF-001-M`

### Goal

Deliver the first usable end-user flows for single prompt execution and interactive chat.

### Description

Implement the CLI entry points that call the shared application services and support both one-off runs and stateful chat sessions.

### Done

- A single prompt can be executed from the CLI.
- Interactive chat can keep context inside the session.
- The commands use application services rather than talking directly to adapters.
- Errors are returned with consistent messages and exit codes.

---

## ISS-002-M-03 Implement `status` command with exit codes and JSON output

Labels: `phase:1`, `priority:p2`, `area:cli`, `type:feature`

Parent: `ISS-002-M`

Depends on: `ISS-002-M-01`, `ISS-003-M-03`

Related requirements: `F-013-M`, `F-016-M`, `F-018-M`, `NF-008-M`

### Goal

Provide script-friendly runtime and environment status information.

### Description

The CLI should expose a status command that reports effective runtime context and supports machine-readable output for automation.

### Done

- A status command exists.
- The command supports machine-readable output.
- Exit codes are usable in scripts.
- Missing data is reported explicitly instead of silently omitted.

---

## ISS-003-M-01 Define config schema, profile model and precedence rules

Labels: `phase:1`, `priority:p1`, `area:config`, `type:architecture`

Parent: `ISS-003-M`

Related requirements: `F-011-M`, `F-012-M`, `NF-011-M`, `NF-014-M`

### Goal

Define how configuration, profiles, and overrides are represented and merged.

### Description

Noctua needs a stable configuration model for defaults, profile-specific values, model aliases, parameter settings, and path overrides.

### Done

- Config structure is defined.
- Profile structure is defined.
- Precedence rules between defaults, profile values, and CLI overrides are documented.
- The model supports path configuration and provider-specific settings.

---

## ISS-003-M-02 Implement local path resolver and app directory bootstrap

Labels: `phase:1`, `priority:p1`, `area:config`, `type:feature`

Parent: `ISS-003-M`

Depends on: `ISS-003-M-01`

Related requirements: `F-017-M`, `NF-002-M`, `NF-011-M`

### Goal

Bootstrap Noctua local directories in a configurable and predictable way.

### Description

The runtime should resolve where config, data, cache, logs, and artifacts live, and initialize these directories without hardcoded assumptions.

### Done

- Local paths can be resolved from config or environment.
- Missing directories can be created safely.
- Config, data, cache, and logs are separated.
- The solution avoids hardcoded paths in the core.

---

## ISS-003-M-03 Implement config loading and local storage abstractions

Labels: `phase:1`, `priority:p1`, `area:config`, `type:feature`

Parent: `ISS-003-M`

Depends on: `ISS-003-M-01`, `ISS-003-M-02`

Related requirements: `F-011-M`, `F-017-M`, `NF-002-M`, `NF-003-M`

### Goal

Create reusable config and local storage access for the application layer.

### Description

The application layer should load config and access local storage through explicit abstractions rather than direct ad hoc file access.

### Done

- Config loading is implemented behind a clear boundary.
- Local storage access is abstracted for reuse.
- Offline local use works without any cloud dependency.
- Failure modes are typed and actionable.

---

## ISS-004-M-01 Define model catalog, catalog entry and installed model metadata

Labels: `phase:1`, `priority:p1`, `area:models`, `type:architecture`

Parent: `ISS-004-M`

Depends on: `ISS-001-M-01`

Related requirements: `F-005-M`, `F-006-M`, `F-007-M`

### Goal

Define one internal metadata model for installed and available models.

### Description

Noctua needs a shared catalog model before it can list, search, install, delete, or update models consistently across providers.

### Done

- Catalog entries are defined.
- Installed-model metadata is defined.
- Neutral and provider-specific metadata are separated.
- Status fields support lifecycle actions later.

---

## ISS-004-M-02 Implement available-model listing and search across configured sources

Labels: `phase:1`, `priority:p1`, `area:models`, `type:feature`

Parent: `ISS-004-M`

Depends on: `ISS-004-M-01`, `ISS-001-M-03`

Related requirements: `F-006-M`, `F-007-M`

### Goal

List and search available models across configured catalog sources.

### Description

The user should be able to discover available models from configured sources without caring about provider-specific listing logic.

### Done

- Available models can be listed.
- Search and filtering work across configured sources.
- Search results include provider context.
- Failure in one source does not make the entire command unusable without explanation.

---

## ISS-004-M-03 Implement installed-model listing and merged catalog presentation

Labels: `phase:1`, `priority:p2`, `area:models`, `type:feature`

Parent: `ISS-004-M`

Depends on: `ISS-004-M-01`, `ISS-003-M-03`

Related requirements: `F-005-M`, `F-006-M`

### Goal

Present installed and available model information consistently.

### Description

The CLI should be able to show installed models and optionally merge this view with available-model metadata to avoid a fragmented user experience.

### Done

- Installed models can be listed.
- The output clearly differentiates installed and available states.
- Provider and version information are visible.
- Missing or stale metadata is reported clearly.

---

## ISS-005-M-01 Wrap the existing download tool behind an install adapter

Labels: `phase:1`, `priority:p1`, `area:models`, `type:integration`

Parent: `ISS-005-M`

Depends on: `ISS-001-M-01`, `ISS-004-M-01`

Related requirements: `F-008-M`, `NF-006-M`

### Goal

Reuse the existing model download capability behind the new architecture.

### Description

The existing download logic should be integrated through a clean install adapter so the new runtime can benefit from prior work without coupling to implementation details.

### Done

- Existing download functionality is wrapped behind a clear adapter boundary.
- The rest of the system does not depend on tool-specific internals.
- Installation inputs and outputs are normalized.
- Failures are surfaced as typed install errors.

---

## ISS-005-M-02 Implement install state machine with staging and atomic promotion

Labels: `phase:1`, `priority:p1`, `area:models`, `type:feature`

Parent: `ISS-005-M`

Depends on: `ISS-005-M-01`, `ISS-003-M-03`

Related requirements: `F-008-M`, `NF-006-M`

### Goal

Prevent broken installations and inconsistent model states.

### Description

Model installation should move through explicit states such as staged, verified, installed, and failed, with cleanup and atomic promotion where possible.

### Done

- Install states are explicit.
- Temporary and final locations are separated.
- Failed installs can be detected and cleaned up.
- Successful installs only become visible after promotion.

---

## ISS-005-M-03 Add CLI install flow and install job reporting

Labels: `phase:1`, `priority:p2`, `area:models`, `type:feature`

Parent: `ISS-005-M`

Depends on: `ISS-002-M-01`, `ISS-005-M-02`

Related requirements: `F-008-M`, `F-016-M`

### Goal

Expose model installation to end users through the CLI.

### Description

Users should be able to install a model by command and receive understandable progress or result reporting.

### Done

- A CLI install command exists.
- Install success and failure are visible.
- Output is usable both for humans and automation.
- The install command integrates with model catalog state.

---

## ISS-006-M-01 Implement model selection resolution for global, profile and call scopes

Labels: `phase:1`, `priority:p1`, `area:runtime`, `type:feature`

Parent: `ISS-006-M`

Depends on: `ISS-001-M-02`, `ISS-003-M-03`, `ISS-004-M-01`

Related requirements: `F-009-M`, `F-013-M`

### Goal

Resolve the effective model target consistently across all supported scopes.

### Description

Noctua should determine the active model from global defaults, profile settings, and per-command overrides in a deterministic way.

### Done

- Global, profile, and per-call selection are supported.
- Precedence is deterministic.
- Resolution errors are actionable.
- The result can be reused by `run`, `chat`, and later automation flows.

---

## ISS-006-M-02 Persist active model mappings and alias targets

Labels: `phase:1`, `priority:p2`, `area:runtime`, `type:feature`

Parent: `ISS-006-M`

Depends on: `ISS-006-M-01`, `ISS-003-M-03`

Related requirements: `F-009-M`, `F-017-M`

### Goal

Persist active model selections and alias mappings locally.

### Description

The runtime should remember durable model mappings and active defaults where appropriate.

### Done

- Active model mappings are persisted.
- Alias targets can be reloaded after restart.
- Invalid persisted targets are detected cleanly.
- Persistence remains local-first.

---

## ISS-006-M-03 Expose runtime context output for model, provider, profile and thinking state

Labels: `phase:1`, `priority:p2`, `area:runtime`, `type:feature`

Parent: `ISS-006-M`

Depends on: `ISS-006-M-01`, `ISS-008-M-02`

Related requirements: `F-013-M`

### Goal

Make the effective runtime context transparent to the user.

### Description

Users should be able to see which model, provider, profile, and thinking-related mode are active for a given request or status output.

### Done

- Runtime context output includes model, provider or engine, and profile.
- Thinking mode status is represented explicitly.
- Missing or inherited values are handled clearly.
- The same resolved context can be used by both status and execution commands.

---

## ISS-007-M-01 Define session and message storage schema

Labels: `phase:1`, `priority:p1`, `area:history`, `type:architecture`

Parent: `ISS-007-M`

Depends on: `ISS-003-M-03`

Related requirements: `F-010-M`, `NF-002-M`, `NF-003-M`

### Goal

Define how chat sessions and messages are stored locally.

### Description

Noctua needs a storage schema for sessions, messages, and related runtime metadata before it can support resume and history safely.

### Done

- Session and message schema are defined.
- Relationships between session, turn, and message are clear.
- The schema supports local resume flows.
- The design can evolve with versioned migrations later.

---

## ISS-007-M-02 Implement chat session persistence and resume service

Labels: `phase:1`, `priority:p1`, `area:history`, `type:feature`

Parent: `ISS-007-M`

Depends on: `ISS-007-M-01`, `ISS-003-M-03`

Related requirements: `F-010-M`

### Goal

Persist chats locally and allow them to be resumed later.

### Description

The application layer should save messages as part of a session and reload them for continued conversation.

### Done

- Sessions can be created and persisted.
- Existing sessions can be loaded.
- A resumed session preserves prior context order.
- Failed writes do not silently corrupt the conversation state.

---

## ISS-007-M-03 Add CLI support for resume, history listing and continue flow

Labels: `phase:1`, `priority:p2`, `area:history`, `type:feature`

Parent: `ISS-007-M`

Depends on: `ISS-002-M-02`, `ISS-007-M-02`

Related requirements: `F-010-M`, `F-016-M`

### Goal

Expose session continuation and history access in the CLI.

### Description

Users should be able to find a saved session and continue it without reconstructing context manually.

### Done

- Saved sessions can be listed.
- A chosen session can be resumed.
- Continue flow works from the CLI.
- Session-related output remains usable for scripts when needed.

---

## ISS-008-M-01 Define AI parameter schema and validation rules

Labels: `phase:1`, `priority:p1`, `area:config`, `type:architecture`

Parent: `ISS-008-M`

Depends on: `ISS-003-M-01`

Related requirements: `F-012-M`

### Goal

Define one consistent parameter model for model execution settings.

### Description

Parameters such as temperature, token limits, and system prompts should be represented consistently across config, profiles, and runtime overrides.

### Done

- Parameter schema is defined.
- Validation rules are defined.
- Unsupported or malformed values fail clearly.
- The model is provider-neutral at the core boundary.

---

## ISS-008-M-02 Implement parameter merge logic for config, profile and CLI overrides

Labels: `phase:1`, `priority:p1`, `area:config`, `type:feature`

Parent: `ISS-008-M`

Depends on: `ISS-008-M-01`, `ISS-003-M-03`

Related requirements: `F-011-M`, `F-012-M`

### Goal

Compute the effective execution parameters deterministically.

### Description

The runtime should merge defaults, profile values, and CLI overrides into a final validated parameter set for each request.

### Done

- Effective parameter calculation is deterministic.
- CLI overrides only affect the current call when intended.
- Parameter validation happens before provider execution.
- The final parameter set can be inspected later.

---

## ISS-008-M-03 Expose parameter override flags and resolved parameter inspection

Labels: `phase:1`, `priority:p2`, `area:config`, `type:feature`

Parent: `ISS-008-M`

Depends on: `ISS-002-M-01`, `ISS-008-M-02`

Related requirements: `F-012-M`, `F-016-M`

### Goal

Make execution parameter overrides visible and user-controllable from the CLI.

### Description

Users should be able to override parameters in commands and inspect the effective resolved values when needed.

### Done

- Relevant parameter override flags exist.
- Resolved parameters can be inspected.
- Output is clear enough for both humans and scripts.
- Invalid overrides fail before execution.

---

## ISS-009-M-01 Define uninstall safety checks and artifact ownership rules

Labels: `phase:1`, `priority:p1`, `area:models`, `type:architecture`

Parent: `ISS-009-M`

Depends on: `ISS-004-M-01`, `ISS-005-M-02`

Related requirements: `F-014-M`, `NF-006-M`

### Goal

Define the safety model for deleting installed models and related artifacts.

### Description

Before implementing deletion, Noctua needs rules for ownership, references, and safety conditions so removal does not break unrelated state.

### Done

- Deletion preconditions are defined.
- Artifact ownership rules are defined.
- Shared or referenced data is handled explicitly.
- Failure and rollback behavior are documented.

---

## ISS-009-M-02 Implement delete workflow with cleanup and failure recovery

Labels: `phase:1`, `priority:p1`, `area:models`, `type:feature`

Parent: `ISS-009-M`

Depends on: `ISS-009-M-01`, `ISS-003-M-03`

Related requirements: `F-014-M`, `NF-006-M`

### Goal

Delete installed models safely and recover from partial failures.

### Description

The runtime should support a controlled delete workflow that cleans up artifacts and preserves a consistent metadata state.

### Done

- Installed models can be deleted through the core workflow.
- Partial failures are detectable.
- Cleanup avoids leaving broken installed-state metadata.
- The system remains consistent after failed delete attempts.

---

## ISS-009-M-03 Add CLI delete command with non-interactive safety flags

Labels: `phase:1`, `priority:p2`, `area:models`, `type:feature`

Parent: `ISS-009-M`

Depends on: `ISS-002-M-01`, `ISS-009-M-02`

Related requirements: `F-014-M`, `F-016-M`

### Goal

Expose model deletion safely to both users and scripts.

### Description

The CLI should allow model deletion with flags that make automation possible without sacrificing safety.

### Done

- A CLI delete command exists.
- Non-interactive flows are supported explicitly.
- Success and failure states are visible.
- The command integrates with installed-model listings.

---

## ISS-010-M-01 Define structured log and event schema plus redaction rules

Labels: `phase:1`, `priority:p1`, `area:observability`, `type:architecture`

Parent: `ISS-010-M`

Depends on: `ISS-001-M-01`, `ISS-003-M-01`

Related requirements: `NF-007-M`, `NF-012-M`

### Goal

Create one structured observability model for logs and runtime events.

### Description

Noctua should emit structured logs and events from the beginning so later debugging, audit, and support do not depend on ad hoc text output.

### Done

- A structured event schema is defined.
- A structured log schema is defined.
- Redaction rules for secrets and sensitive data are defined.
- The model is usable across CLI and model lifecycle flows.

---

## ISS-010-M-02 Implement resource probes for CPU, RAM and GPU status

Labels: `phase:1`, `priority:p2`, `area:observability`, `type:feature`

Parent: `ISS-010-M`

Depends on: `ISS-002-M-03`, `ISS-003-M-02`

Related requirements: `F-018-M`

### Goal

Expose basic resource awareness for runtime diagnosis.

### Description

The runtime should detect CPU and memory usage, and GPU or VRAM when available, without failing noisily on systems where some metrics are unavailable.

### Done

- CPU and RAM probes exist.
- GPU or VRAM probing is supported where available.
- Unsupported metrics are reported explicitly.
- Resource data can feed status output.

---

## ISS-010-M-03 Wire logs, metrics and resource snapshots into CLI and model lifecycle flows

Labels: `phase:1`, `priority:p2`, `area:observability`, `type:feature`

Parent: `ISS-010-M`

Depends on: `ISS-010-M-01`, `ISS-010-M-02`, `ISS-002-M-02`, `ISS-005-M-03`, `ISS-009-M-03`

Related requirements: `F-018-M`, `NF-004-M`, `NF-005-M`, `NF-012-M`

### Goal

Make observability part of real end-user and model lifecycle flows.

### Description

Structured logs, status signals, and resource snapshots should be emitted from execution, installation, and deletion flows so performance and failures become diagnosable.

### Done

- CLI execution flows emit structured events.
- Install and delete flows emit structured events.
- Status output can surface resource snapshots.
- Secrets remain redacted in all emitted data.
