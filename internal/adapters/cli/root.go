package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Coditary/Noctua/internal/adapters/system/paths"
	"github.com/Coditary/Noctua/internal/app/runtime"
	"github.com/Coditary/Noctua/internal/core/domain"
)

type Dependencies struct {
	Writer      io.Writer
	ErrWriter   io.Writer
	Logger      *slog.Logger
	Config      domain.AppConfig
	ConfigPath  string
	Directories paths.Directories
	Registry    *runtime.ProviderRegistry
	Executor    *runtime.Executor
}

type rootOptions struct {
	json    bool
	profile string
}

func NewRootCommand(deps Dependencies) *cobra.Command {
	if deps.Writer == nil {
		deps.Writer = io.Discard
	}
	if deps.ErrWriter == nil {
		deps.ErrWriter = io.Discard
	}

	options := &rootOptions{}

	root := &cobra.Command{
		Use:           "noctua",
		Short:         "Noctua CLI scaffold",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.SetOut(deps.Writer)
	root.SetErr(deps.ErrWriter)
	root.PersistentFlags().BoolVar(&options.json, "json", false, "Render machine-readable output")
	root.PersistentFlags().StringVar(&options.profile, "profile", "", "Profile to use")

	root.AddCommand(newRunCommand(deps, options, "run"))
	root.AddCommand(newRunCommand(deps, options, "chat"))
	root.AddCommand(newStatusCommand(deps, options))
	root.AddCommand(newModelsCommand(deps, options))

	return root
}

func newRunCommand(deps Dependencies, root *rootOptions, mode string) *cobra.Command {
	var provider string
	var model string
	var alias string
	var systemPrompt string
	var temperature float64
	var maxTokens int
	var thinking bool

	command := &cobra.Command{
		Use:  mode + " [prompt]",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			overrides := domain.CommandOverrides{Profile: root.profile}
			if cmd.Flags().Changed("provider") {
				overrides.Provider = &provider
			}
			if cmd.Flags().Changed("model") {
				overrides.Model = &model
			}
			if cmd.Flags().Changed("alias") {
				overrides.Alias = &alias
			}
			if cmd.Flags().Changed("system-prompt") {
				overrides.SystemPrompt = &systemPrompt
			}
			if cmd.Flags().Changed("temperature") {
				overrides.Temperature = &temperature
			}
			if cmd.Flags().Changed("max-tokens") {
				overrides.MaxTokens = &maxTokens
			}
			if cmd.Flags().Changed("thinking") {
				overrides.Thinking = &thinking
			}

			result, err := deps.Executor.RunPrompt(context.Background(), deps.Config, overrides, strings.Join(args, " "))
			if err != nil {
				return err
			}

			payload := map[string]any{
				"mode":     mode,
				"settings": result.Settings,
				"result":   result.Result,
			}

			return writeOutput(deps.Writer, root.json, payload, renderRunText(mode, result))
		},
	}

	command.Flags().StringVar(&provider, "provider", "", "Provider override")
	command.Flags().StringVar(&model, "model", "", "Model override")
	command.Flags().StringVar(&alias, "alias", "", "Model alias override")
	command.Flags().StringVar(&systemPrompt, "system-prompt", "", "System prompt override")
	command.Flags().Float64Var(&temperature, "temperature", 0, "Temperature override")
	command.Flags().IntVar(&maxTokens, "max-tokens", 0, "Max token override")
	command.Flags().BoolVar(&thinking, "thinking", false, "Enable thinking mode")

	return command
}

func newStatusCommand(deps Dependencies, root *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use: "status",
		RunE: func(cmd *cobra.Command, args []string) error {
			settings, err := runtime.ResolveSettings(deps.Config, domain.CommandOverrides{Profile: root.profile})
			if err != nil {
				return err
			}

			payload := map[string]any{
				"config_path": deps.ConfigPath,
				"directories": deps.Directories,
				"providers":   deps.Registry.Names(),
				"settings":    settings,
			}

			text := fmt.Sprintf(
				"config: %s\nprofile: %s\nprovider: %s\nmodel: %s\nproviders: %s\n",
				deps.ConfigPath,
				settings.Profile,
				settings.Provider,
				settings.Model,
				strings.Join(deps.Registry.Names(), ", "),
			)

			return writeOutput(deps.Writer, root.json, payload, text)
		},
	}
}

func newModelsCommand(deps Dependencies, root *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use: "models",
		RunE: func(cmd *cobra.Command, args []string) error {
			models, err := deps.Executor.ListModels(context.Background())
			if err != nil {
				return err
			}

			textLines := make([]string, 0, len(models))
			for _, model := range models {
				textLines = append(textLines, fmt.Sprintf("%s\t%s", model.Provider, model.Name))
			}

			return writeOutput(deps.Writer, root.json, map[string]any{"models": models}, strings.Join(textLines, "\n")+"\n")
		},
	}
}

func writeOutput(writer io.Writer, asJSON bool, payload any, text string) error {
	if asJSON {
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "  ")
		return encoder.Encode(payload)
	}

	_, err := io.WriteString(writer, text)
	return err
}

func renderRunText(mode string, result runtime.RunResult) string {
	return fmt.Sprintf(
		"mode: %s\nprofile: %s\nprovider: %s\nmodel: %s\nresponse: %s\n",
		mode,
		result.Settings.Profile,
		result.Result.Provider.Name,
		result.Result.Model.Name,
		domain.MessageText(result.Result.Message),
	)
}
