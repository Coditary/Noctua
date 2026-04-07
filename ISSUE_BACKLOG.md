# Noctua Issue Backlog

Detailed Phase 1 planning also lives in:

- `PHASE1_ISSUE_BREAKDOWN.md`
- `GITHUB_ISSUES_PHASE1.md`

The child issue drafts in `GITHUB_ISSUES_PHASE1.md` are intentionally copy-paste-ready for later GitHub issue creation.

These entries are written as high-level GitHub issues or epics. Each entry references requirement IDs and can later be broken down into smaller implementation tasks.

## Phase 1 - Must

### ISS-001-M Core runtime and provider abstraction

Labels: `epic`, `phase:1`, `priority:must`, `area:core`

Covers: `F-001-M`, `F-004-M`, `F-015-M`, `NF-009-M`, `NF-010-M`, `NF-014-M`

Description: Define the core abstractions for engine, provider, model alias, runtime context, and execution request. The core must be able to address local engines and external providers through the same flow.

Acceptance Criteria:
- A clear interface exists for local and external execution adapters.
- At least one local and one external provider use the same request path.
- Extension points for new providers are documented and versionable.

### ISS-002-M CLI foundation for run, chat, and status

Labels: `epic`, `phase:1`, `priority:must`, `area:cli`

Covers: `F-002-M`, `F-003-M`, `F-016-M`, `NF-001-M`, `NF-008-M`, `NF-013-M`

Description: Implement the first production-ready CLI with non-interactive run mode, interactive chat, clear exit codes, and script-friendly output.

Acceptance Criteria:
- A single-run prompt works without an interactive session.
- An interactive session retains context across multiple messages.
- Important commands provide both human-readable and machine-readable output.
- Exit codes are stable enough for scripts to evaluate reliably.

### ISS-003-M Local configuration, profiles, and storage layout

Labels: `epic`, `phase:1`, `priority:must`, `area:config`

Covers: `F-011-M`, `F-017-M`, `NF-002-M`, `NF-011-M`, `NF-014-M`

Description: Define the local data model for config, profiles, history, caches, and metadata. Users must be able to define defaults centrally and relocate data paths.

Acceptance Criteria:
- A configuration file supports defaults for model, provider, and parameters.
- Multiple profiles can be loaded and overridden.
- Data, cache, and config paths are configurable.
- Local usage works without cloud storage.

### ISS-004-M Model catalog for installed and available models

Labels: `epic`, `phase:1`, `priority:must`, `area:models`

Covers: `F-005-M`, `F-006-M`, `F-007-M`, `F-004-M`

Description: Build a shared model catalog that represents installed models, available models, and search or filter functions across providers.

Acceptance Criteria:
- Installed models are shown with provider, tag or version, and status.
- Available models can be listed from multiple sources.
- Search and filtering by name, tag, and provider work.

### ISS-005-M Automated model installation and download pipeline

Labels: `epic`, `phase:1`, `priority:must`, `area:models`

Covers: `F-008-M`, `NF-006-M`

Description: Integrate automated model installation into the core. Existing download logic should be reused if it fits the planned architecture.

Acceptance Criteria:
- A model can be installed through the CLI.
- Interrupted or failed downloads leave no inconsistent state behind.
- Installation failures are clearly visible to both users and scripts.

### ISS-006-M Model selection and transparent runtime context

Labels: `epic`, `phase:1`, `priority:must`, `area:runtime`

Covers: `F-009-M`, `F-013-M`

Description: Introduce global, profile-based, and per-call model selection, and show the active runtime context transparently for every request.

Acceptance Criteria:
- A default model can be set globally or per profile.
- A command can override the model only for the current invocation.
- Output shows the active model, provider or engine, profile, and thinking mode status.

### ISS-007-M Persist chat history and resume sessions

Labels: `epic`, `phase:1`, `priority:must`, `area:history`

Covers: `F-010-M`, `NF-002-M`, `NF-003-M`

Description: Persist session histories locally and make them available for later resumption.

Acceptance Criteria:
- A session is stored locally.
- A saved session can be loaded and continued later.
- History works in fully local offline usage.

### ISS-008-M Parameter management and CLI overrides

Labels: `epic`, `phase:1`, `priority:must`, `area:config`

Covers: `F-012-M`, `F-011-M`

Description: Implement central AI parameter management with defaults, profiles, and per-call overrides.

Acceptance Criteria:
- Temperature, max tokens, and system prompt are configurable.
- CLI flags can override config values selectively.
- The final active parameter set is traceable.

### ISS-009-M Model deletion and clean cleanup

Labels: `epic`, `phase:1`, `priority:must`, `area:models`

Covers: `F-014-M`, `NF-006-M`

Description: Installed models must be removable in a controlled way without leaving broken references or stale data in an invalid state.

Acceptance Criteria:
- Models can be deleted through a CLI command.
- After deletion, the model disappears from installed-model listings.
- Orphaned artifacts and broken references are handled cleanly.

### ISS-010-M Resource visibility, logging, and baseline observability

Labels: `epic`, `phase:1`, `priority:must`, `area:observability`

Covers: `F-018-M`, `NF-004-M`, `NF-005-M`, `NF-007-M`, `NF-012-M`

Description: Implement runtime status, resource visibility, and structured logs for debugging, support, and performance measurement.

Acceptance Criteria:
- A status call shows CPU, RAM, and, when available, GPU or VRAM.
- Logs are structured and contain no plaintext secrets.
- CLI latency for non-inference commands is measurable.

## Phase 2 - Should

### ISS-101-S TUI shell for chat, history, and models

Labels: `epic`, `phase:2`, `priority:should`, `area:tui`

Covers: `F-101-S`, `NF-101-S`

Description: Build a TUI that makes the most important CLI core workflows more convenient for day-to-day usage.

Acceptance Criteria:
- Chat, model selection, history, and resource status are reachable in the TUI.
- Core flows require no more than three navigation steps.

### ISS-102-S Guided setup for MCP and standard tools

Labels: `epic`, `phase:2`, `priority:should`, `area:onboarding`

Covers: `F-102-S`, `NF-107-S`

Description: Create a guided setup flow that quickly makes recommended MCP servers, tools, and standard components usable.

Acceptance Criteria:
- Users can activate a recommended toolchain through a guided flow.
- The flow is documented and can be completed quickly on reference hardware.

### ISS-103-S MCP client and server integration

Labels: `epic`, `phase:2`, `priority:should`, `area:mcp`

Covers: `F-103-S`

Description: Integrate MCP so that Noctua can both consume external MCP servers and expose its own MCP functionality.

Acceptance Criteria:
- Noctua can consume at least one MCP server.
- Noctua can expose at least one tool or resource through MCP.

### ISS-104-S Registry for custom commands and tools

Labels: `epic`, `phase:2`, `priority:should`, `area:extensibility`

Covers: `F-104-S`, `F-105-S`, `NF-106-S`

Description: Introduce a registrable extension layer for custom commands and tools, including policy control.

Acceptance Criteria:
- Custom commands can be registered and loaded.
- Custom tools can be integrated into tool-calling flows.
- Policies can allow or forbid specific tools or commands.

### ISS-105-S Agent skills and declarative agent definitions

Labels: `epic`, `phase:2`, `priority:should`, `area:agents`

Covers: `F-106-S`, `F-108-S`

Description: Define the format and runtime for agent skills and declarative agents, for example in Markdown plus structured metadata.

Acceptance Criteria:
- A skill can be loaded and assigned to an agent.
- An agent definition can be registered without code changes.
- The format is documented and extensible for plugins.

### ISS-106-S Multi-agent orchestration and routing

Labels: `epic`, `phase:2`, `priority:should`, `area:agents`

Covers: `F-107-S`, `F-114-S`, `NF-105-S`

Description: Implement the role model of Interpreter, Orchestrator, Specialist, and Evaluator together with routing logic for speed, cost, and availability.

Acceptance Criteria:
- A task can pass through all four roles.
- The flow is logged and traceable.
- Policies can prefer local, remote, or mixed execution.

### ISS-107-S Better tool calling for local models and worker launch

Labels: `epic`, `phase:2`, `priority:should`, `area:tools`

Covers: `F-109-S`, `F-110-S`

Description: Improve tool calling for local models, including structured parameters, guardrails, and controlled startup of external AI workers.

Acceptance Criteria:
- Local models can produce validated tool calls.
- Tool failures are handled robustly and reported back.
- External workers such as image generators can be launched in a controlled way through tools.

### ISS-108-S Server mode, remote client, and SDK or API

Labels: `epic`, `phase:2`, `priority:should`, `area:remote`

Covers: `F-111-S`, `F-112-S`, `NF-102-S`, `NF-103-S`, `NF-104-S`

Description: Expose the core both as a network service and as a local programming interface.

Acceptance Criteria:
- A Noctua server can be started.
- A remote client can send requests to the server.
- A third-party application can call the core through an SDK or API directly.
- Remote connections are secured.
- Versioned data and configuration migrations are planned for upgrades.

### ISS-109-S Model updates and version management

Labels: `epic`, `phase:2`, `priority:should`, `area:models`

Covers: `F-113-S`, `NF-006-M`

Description: Extend the model lifecycle with update checks, version handling, and traceable upgrade paths.

Acceptance Criteria:
- An update command detects new model versions.
- Updates can be run or aborted in a controlled way.
- The old and new states are documented traceably.

## Phase 3 - Could

### ISS-201-C Web app foundation for chat and admin surface

Labels: `epic`, `phase:3`, `priority:could`, `area:web`

Covers: `F-201-C`, `F-204-C`, `NF-201-C`

Description: Build the first web interface for chat, session history, and basic administration.

Acceptance Criteria:
- Chat and history are available on the web.
- Basic model and status information is visible.
- The interface meets baseline accessibility expectations.

### ISS-202-C Auth, user profiles, and isolated histories

Labels: `epic`, `phase:3`, `priority:could`, `area:web`

Covers: `F-202-C`, `F-203-C`, `NF-202-C`

Description: Introduce login, per-user settings, and clear isolation between users in multi-user operation.

Acceptance Criteria:
- Users can sign in and sign out.
- Settings are stored per user.
- Histories and data are strictly isolated.

### ISS-203-C Remote admin client for model management

Labels: `epic`, `phase:3`, `priority:could`, `area:admin`

Covers: `F-205-C`, `NF-203-C`

Description: Build an admin capability or separate client that lets hosted Noctua instances download and manage models remotely.

Acceptance Criteria:
- An admin can trigger remote model installations.
- Critical admin actions are logged in an auditable way.

### ISS-207-C Platform portability for macOS and Windows

Labels: `epic`, `phase:3`, `priority:could`, `area:platform`

Covers: `NF-204-C`

Description: Extend the stabilized CLI core from Phases 1 and 2 to additional target platforms with reproducible core flows and documented platform-specific behavior.

Acceptance Criteria:
- The defined core flows run reproducibly on macOS and Windows.
- Platform-specific behavior for paths, processes, and resource inspection is documented.
- Platform-independent tests or checks cover the most important differences.

### ISS-204-C Speech stack for TTS, STT, and voice management

Labels: `epic`, `phase:3`, `priority:could`, `area:speech`

Covers: `F-206-C`, `F-207-C`, `F-208-C`, `F-209-C`

Description: Integrate speech input and output together with later voice management and voice training.

Acceptance Criteria:
- Text can be converted to speech.
- Speech can be transcribed to text.
- Voices can be switched and later trained or adapted.

### ISS-205-C Training and fine-tuning workflows

Labels: `epic`, `phase:3`, `priority:could`, `area:training`

Covers: `F-210-C`, `F-211-C`, `NF-205-C`

Description: Build guided workflows for fine-tuning, dataset management, and resumable training jobs.

Acceptance Criteria:
- A fine-tuning job can be started through a guided flow.
- Files or datasets can be assigned to jobs.
- Long-running jobs can be observed and resumed.

### ISS-206-C Multimodal modules for image, video, audio, 3D, documents, and retrieval

Labels: `epic`, `phase:3`, `priority:could`, `area:multimodal`

Covers: `F-212-C`, `F-213-C`, `F-214-C`, `F-215-C`, `F-216-C`, `F-217-C`, `F-218-C`, `F-219-C`

Description: Extend Noctua beyond text into multimodal and structured AI tasks. The architecture should allow new categories to be added later without core redesign.

Acceptance Criteria:
- At least one image flow, one audio or speech flow, and one document or retrieval flow can be integrated.
- New modalities use defined adapters instead of special-case logic in the core.
- Structured tasks such as translation or summarization can run through the same routing layer.
