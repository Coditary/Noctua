# Noctua

Product planning documents live in:

- `REQUIREMENTS.md`
- `ISSUE_BACKLOG.md`
- `PHASE1_ISSUE_BREAKDOWN.md`
- `PHASE2_ISSUE_BREAKDOWN.md`
- `GITHUB_ISSUES_PHASE1.md`
- `ARCHITECTURE.md`

`GITHUB_ISSUES_PHASE1.md` already contains copy-paste-ready issue drafts for the initial Phase 1 backlog.

## Current Scaffold

The repository now also contains the first Go-based CLI scaffold for Phase 1 issue `#11`.

- entrypoint: `cmd/noctua/main.go`
- core contracts: `internal/core/`
- runtime and registry scaffolding: `internal/app/runtime/`
- CLI adapter: `internal/adapters/cli/`
- config and path bootstrap: `internal/adapters/config/`, `internal/adapters/system/`
- stub providers for local and remote execution: `internal/adapters/providers/stub/`

Quick verification commands:

- `go test ./...`
- `go run ./cmd/noctua status --json`
