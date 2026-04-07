# Noctua Architecture

## Purpose

This document translates the product vision from `REQUIREMENTS.md` into a practical target architecture. The central delivery rule remains unchanged: first a stable CLI core, then a TUI, then a web layer. Noctua should unify local engines, external providers, tools, MCP, and later agents through one common execution model, without each interface bringing its own business logic.

## Architecture Style

Noctua should start as a modular monolith with a hexagonal core.

Why this shape:

1. A single core can be reused by the CLI first, then the TUI, and later the web layer.
2. Local and offline-capable operation is easier than with early distributed services.
3. Providers, engines, tools, and MCP can be separated cleanly through ports and adapters.
4. Extensibility remains possible without overcomplicating the Phase 1 MVP.

## Design Principles

1. CLI-first, TUI-second, Web-third.
2. Core-first before UI-first.
3. Engine, provider, model, alias, tool, and agent are separate concepts.
4. Local storage is the default.
5. New modalities are added through adapters, not core rewrites.
6. UI layers call only application services, never providers or storage directly.
7. Tool and remote actions must be policy-controlled and observable.

## Layers

### 1. Interface Layer

Responsible for input, output, and user-facing surfaces.

- CLI adapter
- TUI adapter
- Web or API adapter
- SDK or library adapter
- MCP boundary for client use and later server exposure

This layer contains no business logic. It maps inputs to use cases and presents results or errors.

### 2. Application Layer

Responsible for concrete use cases.

Examples of services or use cases:

- `run prompt`
- `chat start`
- `chat continue`
- `models list`
- `models search`
- `models install`
- `models delete`
- `profiles resolve`
- `status show`

This layer validates inputs, builds a runtime context, calls routers or services, and returns stable results plus exit codes.

### 3. Domain Layer

Responsible for stable business concepts, independent of UI and implementation technology.

Key models:

- `ProviderRef`
- `EngineRef`
- `ModelRef`
- `ModelAlias`
- `CapabilitySet`
- `Profile`
- `ConfigSnapshot`
- `ExecutionRequest`
- `ExecutionContext`
- `ExecutionResult`
- `Session`
- `Message`
- `ContentPart`
- `CatalogEntry`
- `InstalledModel`
- `InstallJob`
- `ToolSpec`
- `ToolCall`
- `ToolResult`
- `AgentDefinition`
- `SkillDefinition`
- `JobState`
- `ResourceSnapshot`
- `AuditEvent`

Important for later multimodality: `ContentPart` should support more than plain text from the beginning, even if Phase 1 only uses text and chat.

### 4. Orchestration Layer

Responsible for more complex execution flows above individual use cases.

Building blocks:

- Request router
- Model resolver
- Execution planner
- Tool runtime coordinator
- MCP coordinator
- Agent orchestrator
- Evaluator loop

The later role model `Interpreter -> Orchestrator -> Specialist -> Evaluator` should be implemented as one shared runtime with different role profiles, not as four isolated systems.

### 5. Adapter Layer

Responsible for concrete technical integrations.

- Local engine adapters, for example Ollama or later llama.cpp
- Remote provider adapters, for example OpenAI-compatible APIs
- Model catalog source adapters
- Download and install adapters
- Tool runtime adapters
- MCP client adapters
- MCP server adapters
- Secret adapters
- Resource metrics adapters

All adapters should speak inward only through stable ports from the core and application layers.

### 6. Persistence Layer

Responsible for local persistence and later server-side persistence.

Recommendation:

- SQLite for metadata, sessions, models, jobs, events, and policies
- Filesystem for larger artifacts, exports, temporary downloads, and later training data
- separate secret handling for API keys and similar values

Why not files only: file-based storage is fine for individual configs, but too fragile for history, jobs, status, and future migrations. Why not database only: model artifacts, exports, and larger binary files still belong in the filesystem.

### 7. Platform Services

Cross-cutting services for the whole runtime.

- Logging
- Eventing
- Caching
- Process abstraction
- File abstraction
- Resource inspection
- Secret resolution
- Download verification
- Cleanup and retry strategies

## Architecture Overview

```text
Users / Scripts / Future Clients
        |
        v
CLI / TUI / Web / SDK / MCP boundary
        |
        v
Application Services
        |
        v
Runtime Router / Coordinators / Use Case Logic
   |            |              |
   v            v              v
Provider     Tool Runtime     MCP
Adapters     Adapters         Adapters
   |            |              |
   +------------+--------------+
                |
                v
Persistence + Platform Services
```

## Concept Boundaries

These distinctions must remain strict in the core:

- Engine: a local runtime such as Ollama or later llama.cpp.
- Provider: an execution service or endpoint, local or remote.
- Model: a concrete model selectable on an engine or provider.
- Alias: a user-stable reference that can later be remapped to another target.
- Tool: an executable action outside model inference.
- Agent: a defined workflow or role profile running on the same runtime.

## Phase 1 Request Flow

A typical Phase 1 flow for `run` or `chat` looks like this:

1. The CLI reads profile, config, paths, and overrides.
2. The model resolver resolves an alias or default model to a concrete target.
3. The application layer builds an `ExecutionContext` with model, provider, parameters, and session state.
4. The matching provider or engine adapter executes the request.
5. Result, usage, errors, and session data are persisted locally.
6. Logging and resource status are recorded as events or snapshots.
7. The CLI presents response, status, and exit code.

## Persistence Model

### Local Directory Structure

Recommended categories:

- `config/` for global config and profiles
- `data/` for SQLite, sessions, and metadata
- `cache/` for synchronized catalogs and temporary data
- `logs/` for structured runtime logs
- `artifacts/` for exports, later training data, and other files

Concrete paths should remain configurable.

### What belongs in SQLite

- Profiles and effective configuration snapshots
- Model catalog metadata
- Installed models and alias mappings
- Sessions, messages, and run records
- Jobs for install, update, delete, and later long-running work
- Audit events and structured event logs

### What belongs in the filesystem

- Downloaded artifacts
- Temporary download directories
- Export files
- Later datasets, audio, images, videos, and other larger binary files

## Extension Points

### Provider adapter

Contract for external APIs and local provider frontends.

Minimum responsibilities:

- report capabilities
- list or resolve models
- execute requests
- map errors and usage cleanly
- optionally support streaming

### Engine adapter

Contract for local runtimes.

Minimum responsibilities:

- detect locally installed models
- perform or delegate install, delete, and update
- start inference
- provide resource hints

### Model catalog source

Contract for catalogs and model sources.

- list models
- search models
- resolve versions
- later provide checksums or provenance

### Tool runtime adapter

Contract for custom tools and workers.

- `ToolSpec`
- parameter schema
- invocation
- timeout and retry
- policy check
- audit event

### MCP bridge

Two directions:

- MCP client for external MCP servers and tools
- MCP server for later exposure of Noctua functionality

### Agent definition loader

Recommendation for Phase 2:

- Markdown for readable authoring
- YAML frontmatter for structured metadata
- internal validation into an `AgentDefinition` object

Lua should only be considered later for advanced hooks or extensions, not as the mandatory foundation of the MVP.

### Policy and routing engine

Responsible for decisions such as:

- local vs remote
- fast vs cost-efficient
- offline-only
- tool allowed or forbidden
- provider allowed or forbidden

## Definition Formats

Recommended order:

1. Phase 1: configuration and profiles in a static human-readable format, for example TOML or YAML.
2. Phase 2: agents and skills in Markdown plus YAML frontmatter.
3. Phase 3: optional script hooks only where purely declarative definitions are no longer enough.

This keeps the foundation auditable, reproducible, and safer than allowing arbitrary scripts too early.

## Transport for Later Server Mode

Recommendation for Phase 2 or 3:

- HTTP plus JSON for standard requests
- streaming via Server-Sent Events or later WebSockets

This is simpler for CLI, web, and SDK than introducing a more complex RPC protocol too early.

## Recommended Implementation Order

### Phase 1

1. Core contracts and provider ports
2. Config, profiles, and local storage layout
3. Model catalog and alias resolution
4. Install and delete lifecycle
5. CLI for `run`, `chat`, `status`, and `models`
6. History and session continuation
7. Parameter management and transparent runtime contexts
8. Logging, resource visibility, and observability

### Phase 2

1. TUI on the same application services
2. Tool runtime and MCP client
3. Declarative agents and skills
4. Routing layer and multi-agent roles
5. Server mode and SDK

### Phase 3

1. Web with auth and multi-user context
2. Admin flows and audit
3. Speech, training, and multimodal adapters
4. Later long-running and resumable jobs

## Risks and Guardrails

1. Do not try to build a universal AI operating system too early. Phase 1 must deliver stable text and chat first.
2. Engine, provider, and model must not be mixed together.
3. UI-specific logic must not leak into the core.
4. Tool calling needs guardrails, audit, and policies before real side effects are allowed.
5. Install, update, and delete must run through controlled job states so inconsistent model states cannot occur.
6. Hidden online dependencies would undermine local-first and offline-capable operation.
7. Using Lua too early as the main configuration foundation would make debugging, security, and reproducibility harder.

## Open ADRs

1. Which engines and providers are mandatory in Phase 1: Ollama plus an OpenAI-compatible remote provider, or llama.cpp as well?
2. Which configuration format should Phase 1 use concretely?
3. Is SQLite acceptable as the default persistence mechanism?
4. How should thinking be presented in the product: mode only, or actual reasoning output?
5. What security boundaries should later apply to tools, workers, and MCP?
6. At what point is a hardened server mode required?

## Mapping to Requirements and Epics

| Architecture Building Block | Main Reference |
| --- | --- |
| Provider and engine ports | `F-001-M`, `F-004-M`, `F-015-M`, `ISS-001-M` |
| CLI application services | `F-002-M`, `F-003-M`, `F-016-M`, `ISS-002-M` |
| Config and local storage | `F-011-M`, `F-017-M`, `ISS-003-M` |
| Model catalog and lifecycle | `F-005-M`, `F-006-M`, `F-007-M`, `F-014-M`, `ISS-004-M`, `ISS-005-M`, `ISS-009-M` |
| Model selection and runtime context | `F-009-M`, `F-013-M`, `ISS-006-M`, `ISS-008-M` |
| Session and history | `F-010-M`, `ISS-007-M` |
| Resources and observability | `F-018-M`, `ISS-010-M`, `NF-004-M`, `NF-012-M` |
| MCP, tools, and agents | `F-102-S` through `F-114-S` |
| Web, auth, and multi-user | `F-201-C` through `F-219-C`, `NF-201-C` through `NF-205-C` |
