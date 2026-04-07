# Noctua Requirements

## Product Summary

Noctua is a modular AI platform that unifies local engines, external providers, tools, and agents under one shared runtime. The delivery order is intentionally CLI first, then TUI, then Web. The core should be local-first, offline-capable, and extensible through a plugin or adapter model. Models, engines, providers, tools, and agents should all be addressable through one common execution model. Phase 1 focuses on text and chat, Phase 2 introduces TUI, MCP, and agent orchestration, and Phase 3 expands into Web, multi-user support, and multimodal workflows. Additional useful AI use cases include document understanding, OCR and PDF chat, embeddings and retrieval, image editing, audio transformation, video understanding, and structured tasks such as translation, summarization, and classification.

## Guardrails

1. CLI-first, TUI-second, Web-third.
2. Core-first before UI-first.
3. Engine, provider, model, tool, and agent are separate concepts.
4. Local storage is the default.
5. New modalities must be addable without rebuilding the core.
6. Local usage must not depend on cloud accounts.

## Functional Requirements

### Phase 1 - Must

| ID | Area | Description | Acceptance Criterion |
| --- | --- | --- | --- |
| F-001-M | Core | The core unifies local engines and external providers through a common execution interface. | At least one local provider and one external provider can be called through the same CLI workflow. |
| F-002-M | CLI | The system provides a non-interactive CLI mode for single invocations. | A prompt can be executed by command and returns a response plus a clear exit code. |
| F-003-M | CLI | The system provides an interactive CLI chat mode with context retention within the session. | Multiple inputs in one session correctly reference previous messages. |
| F-004-M | Models | Users can map model aliases to engines, providers, or remote endpoints. | An alias can be configured and later remapped to another provider without changing prompts. |
| F-005-M | Models | The system lists locally installed models per engine or provider. | A CLI command shows name, tag or version, provider, and installation status. |
| F-006-M | Models | The system lists available models from configured sources. | A CLI command can show a unified catalog from multiple sources. |
| F-007-M | Models | Users can search and filter available models. | Search by name, tags, or provider returns filtered results. |
| F-008-M | Models | Models can be downloaded and installed automatically. | A user can install a model through the CLI without manual file handling. |
| F-009-M | Models | Users can switch the active model globally, per profile, or per command. | A CLI override changes the default only for the current invocation. |
| F-010-M | History | Chat histories are stored locally and can be reloaded. | A session can be reopened and continued after it ends. |
| F-011-M | Configuration | The system supports defaults and profiles through a configuration file. | Default values for model, provider, parameters, and paths can be defined centrally. |
| F-012-M | Configuration | Core AI parameters such as temperature, max tokens, and system prompt are configurable. | Parameters can be set through config and through CLI flags. |
| F-013-M | Transparency | CLI output shows the active model, engine or provider, profile, and thinking mode status. | Every response or status call includes these values or explicitly marks them as missing. |
| F-014-M | Models | Installed models can be deleted. | A model can be removed through the CLI and no longer appears as installed afterward. |
| F-015-M | Extensibility | The core provides a stable plugin or adapter model for providers and model execution. | A new provider can be added without rewriting the CLI command model. |
| F-016-M | Automation | The CLI provides script-friendly output formats for automation. | Relevant commands can return machine-readable output in addition to text output. |
| F-017-M | Local Data | Runtime metadata, configuration, and history are trackable through local data structures. | Noctua can resolve the local data associated with a session without cloud dependencies. |
| F-018-M | Resources | The system shows CPU, RAM, and, when available, GPU or VRAM utilization. | A status call shows available resources or clearly reports when metrics are unavailable. |

### Phase 2 - Should

| ID | Area | Description | Acceptance Criterion |
| --- | --- | --- | --- |
| F-101-S | TUI | The system provides a TUI for chat, model switching, history, and resource views. | Core tasks are reachable from the TUI without falling back to raw CLI usage. |
| F-102-S | Onboarding | A guided setup flow installs MCP and standard tools as easily as possible. | A user can activate the recommended standard toolchain through a guided flow. |
| F-103-S | MCP | The system integrates MCP client and MCP server functionality into the runtime. | Noctua can consume at least one MCP server and expose tools itself. |
| F-104-S | Extensibility | Users can define and register their own commands. | A custom command can be made available in the runtime without core changes. |
| F-105-S | Extensibility | Users can define their own tools and use them in tool-calling flows. | A custom tool can be registered, invoked, and logged. |
| F-106-S | Agents | Agent skills can be loaded and assigned to an agent or workflow. | A skill can be referenced and measurably changes an agent's behavior. |
| F-107-S | Agents | The system supports a multi-agent model with Interpreter, Orchestrator, Specialist, and Evaluator. | A task can be handed through these roles and the flow is logged. |
| F-108-S | Agents | Agents or parts of their definitions can be declared declaratively, for example in Markdown plus metadata. | An agent definition can be loaded and executed without code changes. |
| F-109-S | Tool Calling | Tool calling for local models supports structured schemas, guardrails, and error handling. | A local model can call a tool with validated arguments and receive clean error feedback. |
| F-110-S | Worker Control | Noctua can use approved tools to start external AI programs or workers, for example for image generation. | A tool can start an external worker in a controlled way and document the return channel. |
| F-111-S | Remote Operation | Noctua can run as a server and allow external clients to connect. | A remote client can exchange a request with a running Noctua server. |
| F-112-S | Integration | The core is usable as a library or SDK and as an addressable app interface. | A third-party application can call Noctua functionality directly without shell wrapping. |
| F-113-S | Models | Installed models can be updated. | An update command detects new versions, performs the update, and documents the result. |
| F-114-S | Routing | A top-level routing layer can choose the best model or tool path depending on task, cost, speed, availability, and policies. | The same task can run locally, remotely, or multimodally under different policies. |

### Phase 3 - Could

| ID | Area | Description | Acceptance Criterion |
| --- | --- | --- | --- |
| F-201-C | Web | The system provides a web app for chat, administration, and observation. | Chat, model status, and basic administration are available through the web. |
| F-202-C | Web | The web app supports authentication and login. | Users can sign in and only access the areas they are allowed to use. |
| F-203-C | Web | Users have individual settings in the web interface. | Settings are stored per user and restored on the next login. |
| F-204-C | Web | Histories are presented as chats in the web interface and can be resumed. | A user can select an earlier web chat and continue it. |
| F-205-C | Administration | An admin client can control remote model downloads and model availability for hosted environments. | An admin can install a model on a server from a separate client. |
| F-206-C | Speech | The system supports TTS output. | A text response can be converted to audio and played back or exported. |
| F-207-C | Speech | The system supports STT input. | Voice input can be transcribed to text and forwarded to chat or commands. |
| F-208-C | Speech | TTS voices can be changed per user, profile, or session. | A user can select a voice and replay the same request using a different voice. |
| F-209-C | Speech | TTS voices can be trained or adapted. | A defined workflow produces a usable new or adapted voice. |
| F-210-C | Training | Fine-tuning workflows can be launched through guided flows. | A user can trigger fine-tuning without manually orchestrating a pipeline. |
| F-211-C | Training | Files or datasets for training and adaptation jobs can be stored and referenced. | A training job can reference a managed local dataset. |
| F-212-C | Multimodal | Image generation and image editing are supported through providers or tools. | An image workflow can be executed from a prompt. |
| F-213-C | Multimodal | Video generation and video understanding are supported through providers or tools. | A video workflow can be started and its result managed. |
| F-214-C | Multimodal | Audio generation or audio transformation is supported. | An audio workflow can run from text or audio input. |
| F-215-C | Multimodal | 3D generation or 3D modeling workflows are supported. | A 3D workflow can be orchestrated by Noctua and its result referenced. |
| F-216-C | Documents | Document understanding, OCR, PDF chat, and visual document retrieval are supported. | A user can access PDF or document content through chat or retrieval. |
| F-217-C | Knowledge | Embeddings, semantic search, and local retrieval or RAG are supported. | Documents can be indexed and later queried through semantic search. |
| F-218-C | Task Classes | Structured AI tasks such as translation, summarization, classification, or ranking run through the same platform. | A non-generative or semi-generative task can run through the general routing layer. |
| F-219-C | Extensibility | New modality classes or tool classes can be added without redesigning the core. | A new adapter for a new AI category can be integrated through defined extension points. |

## Non-Functional Requirements

### Phase 1 - Must

| ID | Area | Description | Measurable or Verifiable |
| --- | --- | --- | --- |
| NF-001-M | Delivery Order | The delivery order is binding: CLI must be complete before TUI, and TUI before Web. | Release planning and backlog order respect this sequence. |
| NF-002-M | Data Storage | The product is local-first; chats, configuration, logs, and metadata are stored locally by default. | A local installation can operate without cloud storage. |
| NF-003-M | Offline | Core operation must work offline as long as a local provider and local model are available. | Text and chat workflows work without internet access when using a local engine. |
| NF-004-M | Performance | Non-inference CLI commands return in under 500 ms median time on reference hardware. | Measurement on a defined reference machine confirms the target value. |
| NF-005-M | Performance | Request dispatch and provider resolution add no more than 300 ms overhead excluding model runtime. | Benchmarking shows the core is not the primary bottleneck. |
| NF-006-M | Consistency | Model download, update, and delete operations are verifiable, fault-tolerant, and leave no inconsistent state behind. | Interrupted or failed operations can be detected and cleaned up safely. |
| NF-007-M | Security | Secrets such as API keys must never be logged in plain text. | Logs, crash reports, and status output never contain plaintext secrets. |
| NF-008-M | Error Handling | Errors must be reported with understandable messages and machine-readable exit codes. | Scripts can distinguish failure modes and users receive actionable guidance. |
| NF-009-M | Portability | The core contains no hardcoded paths or unnecessary engine-specific assumptions. | Paths, caches, and binaries are abstracted through configuration or adapters. |
| NF-010-M | Extensibility | Extension points for providers, tools, and agents are versioned and documented. | An external developer can build an adapter against a documented interface. |
| NF-011-M | Configuration | Data, cache, and configuration paths are configurable. | A user can relocate local directories through config or environment variables. |
| NF-012-M | Observability | Noctua emits structured logs for executions, providers, tools, and errors. | Core events can be traced for debugging and support. |
| NF-013-M | Independence | Local CLI usage requires no vendor or cloud account. | A user can work purely locally without signing into a third party. |
| NF-014-M | Privacy | Network access only occurs to user-configured providers, registries, or servers. | No unexpected outbound connections occur in normal operation. |

### Phase 2 - Should

| ID | Area | Description | Measurable or Verifiable |
| --- | --- | --- | --- |
| NF-101-S | Usability | The TUI should make core tasks such as chat, model switching, and history access reachable in no more than three navigation steps. | Usability validation confirms the target for core flows. |
| NF-102-S | Security | Server and client connections should be encrypted and authenticated. | Remote access does not transmit sensitive content unencrypted. |
| NF-103-S | Migration | Configuration and data schema migrations should be possible without data loss. | Upgrading to a new version cleanly migrates existing local data. |
| NF-104-S | Stability | SDK and library interfaces should be semantically versioned and remain stable within a minor version. | Integrations do not break on compatible minor upgrades. |
| NF-105-S | Traceability | Tool and agent executions should be auditable and reproducible for debugging. | A run can be reconstructed from inputs, outputs, and tool invocations. |
| NF-106-S | Policies | Remote, tool, and admin actions should be restrictable through policies. | An operator can centrally allow or forbid specific actions. |
| NF-107-S | Onboarding | The recommended standard toolchain should install quickly on reference hardware. | A guided setup can be completed in under ten minutes when dependencies are available. |

### Phase 3 - Could

| ID | Area | Description | Measurable or Verifiable |
| --- | --- | --- | --- |
| NF-201-C | Accessibility | The web app should meet baseline accessibility requirements. | Keyboard usability and sufficient contrast are verified. |
| NF-202-C | Multi-User | User histories and settings must be strictly separated in multi-user operation. | One user cannot access another user's data. |
| NF-203-C | Audit | Admin actions in server or web operation should be auditably logged. | Critical remote actions are traceable with time, user, and result. |
| NF-204-C | Platforms | Additional platform support, especially Windows and macOS, should be added after the CLI core is stable. | Core flows run reproducibly on the additional target platforms. |
| NF-205-C | Long-Running Jobs | Training, fine-tuning, and larger multimodal jobs should be resumable and observable. | Interrupted jobs can be resumed or stopped in a controlled way. |

## Open Product Decisions

1. Which providers and engines are mandatory in Phase 1: only Ollama plus an OpenAI-compatible remote provider, or llama.cpp as well?
2. What should the initial plugin and definition format be: Lua, Markdown plus YAML, JSON, or a combination?
3. Does "show thinking" mean only showing a mode flag, or also exposing actual reasoning content?
4. Which local persistence approach is preferred: file-based, SQLite, or a hybrid?
5. Should Noctua only orchestrate training jobs or later provide a full training pipeline itself?
6. Is secure remote operation starting in Phase 2 sufficient, or is a hardened server variant needed earlier?

## Recommended Phases

### Phase 1 - CLI-first MVP

Includes F-001-M through F-018-M and NF-001-M through NF-014-M. The focus is a stable local and remote-capable text and chat core with model management, configuration, history, and resource visibility.

### Phase 2 - TUI and Automation

Includes F-101-S through F-114-S and NF-101-S through NF-107-S. The focus is better usability, MCP integration, agent orchestration, tool extensibility, and server and SDK usage.

### Phase 3 - Web and Multimodal Platform

Includes F-201-C through F-219-C and NF-201-C through NF-205-C. The focus is multi-user web usage, speech, training, and broader multimodal AI workflows.
