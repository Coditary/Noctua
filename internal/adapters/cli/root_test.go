package cli

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/Coditary/Noctua/internal/adapters/providers/stub"
	"github.com/Coditary/Noctua/internal/adapters/system/paths"
	"github.com/Coditary/Noctua/internal/app/runtime"
	"github.com/Coditary/Noctua/internal/core/domain"
)

func TestStatusCommandJSON(t *testing.T) {
	command, output := testCommand(t)
	command.SetArgs([]string{"status", "--json"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	settings, ok := payload["settings"].(map[string]any)
	if !ok {
		t.Fatal("settings payload missing")
	}
	if settings["provider"] != "local-stub" {
		t.Fatalf("provider = %v, want %q", settings["provider"], "local-stub")
	}
}

func TestRunCommandJSON(t *testing.T) {
	command, output := testCommand(t)
	command.SetArgs([]string{"run", "hello world", "--json"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if !strings.Contains(output.String(), "hello world") {
		t.Fatalf("output = %q, want prompt echoed in stub response", output.String())
	}
}

func testCommand(t *testing.T) (*cobra.Command, *bytes.Buffer) {
	t.Helper()

	buffer := &bytes.Buffer{}
	registry := runtime.NewProviderRegistry()
	if err := registry.Register(stub.NewLocal()); err != nil {
		t.Fatalf("Register(local) error = %v", err)
	}
	if err := registry.Register(stub.NewRemote()); err != nil {
		t.Fatalf("Register(remote) error = %v", err)
	}

	command := NewRootCommand(Dependencies{
		Writer:     buffer,
		ErrWriter:  buffer,
		Logger:     slog.New(slog.NewTextHandler(buffer, nil)),
		Config:     domain.DefaultConfig(),
		ConfigPath: "/tmp/noctua/config.toml",
		Directories: paths.Directories{
			ConfigDir: "/tmp/noctua/config",
			DataDir:   "/tmp/noctua/data",
			CacheDir:  "/tmp/noctua/cache",
			LogDir:    "/tmp/noctua/logs",
		},
		Registry: registry,
		Executor: runtime.NewExecutor(registry),
	})

	return command, buffer
}
