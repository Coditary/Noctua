package bootstrap

import (
	"context"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/Coditary/Noctua/internal/adapters/cli"
	tomlstore "github.com/Coditary/Noctua/internal/adapters/config/toml"
	"github.com/Coditary/Noctua/internal/adapters/observability/logging"
	"github.com/Coditary/Noctua/internal/adapters/providers/stub"
	"github.com/Coditary/Noctua/internal/adapters/system/paths"
	"github.com/Coditary/Noctua/internal/app/runtime"
	"github.com/Coditary/Noctua/internal/core/domain"
)

type App struct {
	root *cobra.Command
}

func New() (*App, error) {
	logger := logging.New(os.Stderr, false)

	resolver := paths.NewResolver("noctua")
	directories, err := resolver.Resolve(domain.PathsConfig{})
	if err != nil {
		return nil, err
	}
	if err := directories.Bootstrap(); err != nil {
		return nil, err
	}

	configPath := filepath.Join(directories.ConfigDir, "config.toml")
	store := tomlstore.New(configPath)
	config, err := store.LoadOptional(context.Background())
	if err != nil {
		return nil, err
	}

	registry := runtime.NewProviderRegistry()
	if err := registry.Register(stub.NewLocal()); err != nil {
		return nil, err
	}
	if err := registry.Register(stub.NewRemote()); err != nil {
		return nil, err
	}

	executor := runtime.NewExecutor(registry)
	root := cli.NewRootCommand(cli.Dependencies{
		Writer:      os.Stdout,
		ErrWriter:   os.Stderr,
		Logger:      logger,
		Config:      config,
		ConfigPath:  configPath,
		Directories: directories,
		Registry:    registry,
		Executor:    executor,
	})

	return &App{root: root}, nil
}

func (a *App) Run(args []string) error {
	a.root.SetArgs(args)
	return a.root.Execute()
}
