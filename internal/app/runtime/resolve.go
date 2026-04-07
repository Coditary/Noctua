package runtime

import (
	"github.com/Coditary/Noctua/internal/core/domain"
	coreerrors "github.com/Coditary/Noctua/internal/core/errors"
)

func ResolveSettings(config domain.AppConfig, overrides domain.CommandOverrides) (domain.ResolvedSettings, error) {
	resolved := domain.ResolvedSettings{
		Profile:      overrides.Profile,
		Provider:     config.Defaults.Provider,
		Model:        config.Defaults.Model,
		Alias:        config.Defaults.Alias,
		SystemPrompt: config.Defaults.SystemPrompt,
		Parameters: domain.ExecutionParameters{
			Temperature: config.Defaults.Temperature,
			MaxTokens:   config.Defaults.MaxTokens,
			Thinking:    config.Defaults.Thinking,
		},
	}

	if resolved.Profile == "" {
		resolved.Profile = config.DefaultProfile
	}
	if resolved.Profile == "" {
		resolved.Profile = domain.DefaultConfig().DefaultProfile
	}

	if resolved.Provider == "" {
		resolved.Provider = domain.DefaultConfig().Defaults.Provider
	}
	if resolved.Model == "" {
		resolved.Model = domain.DefaultConfig().Defaults.Model
	}
	if resolved.Parameters.MaxTokens == 0 {
		resolved.Parameters.MaxTokens = domain.DefaultConfig().Defaults.MaxTokens
	}

	if resolved.Profile != "" {
		profile, ok := config.Profiles[resolved.Profile]
		if !ok && overrides.Profile != "" {
			return domain.ResolvedSettings{}, coreerrors.New(coreerrors.KindConfiguration, "requested profile does not exist")
		}

		applyProfile(&resolved, profile)
	}

	applyOverrides(&resolved, overrides)

	if resolved.Alias != "" {
		alias, ok := config.Aliases[resolved.Alias]
		if !ok {
			return domain.ResolvedSettings{}, coreerrors.New(coreerrors.KindNotFound, "configured alias does not exist")
		}

		if resolved.Provider == "" || resolved.Provider == domain.DefaultConfig().Defaults.Provider {
			resolved.Provider = alias.Provider
		}
		if resolved.Model == "" || resolved.Model == domain.DefaultConfig().Defaults.Model {
			resolved.Model = alias.Model
		}
	}

	if resolved.Provider == "" {
		resolved.Provider = domain.DefaultConfig().Defaults.Provider
	}
	if resolved.Model == "" {
		resolved.Model = domain.DefaultConfig().Defaults.Model
	}
	if resolved.Parameters.MaxTokens == 0 {
		resolved.Parameters.MaxTokens = domain.DefaultConfig().Defaults.MaxTokens
	}

	return resolved, nil
}

func applyProfile(resolved *domain.ResolvedSettings, profile domain.ProfileConfig) {
	if profile.Provider != nil {
		resolved.Provider = *profile.Provider
	}
	if profile.Model != nil {
		resolved.Model = *profile.Model
	}
	if profile.Alias != nil {
		resolved.Alias = *profile.Alias
	}
	if profile.SystemPrompt != nil {
		resolved.SystemPrompt = *profile.SystemPrompt
	}
	if profile.Temperature != nil {
		resolved.Parameters.Temperature = *profile.Temperature
	}
	if profile.MaxTokens != nil {
		resolved.Parameters.MaxTokens = *profile.MaxTokens
	}
	if profile.Thinking != nil {
		resolved.Parameters.Thinking = *profile.Thinking
	}
}

func applyOverrides(resolved *domain.ResolvedSettings, overrides domain.CommandOverrides) {
	if overrides.Provider != nil {
		resolved.Provider = *overrides.Provider
	}
	if overrides.Model != nil {
		resolved.Model = *overrides.Model
	}
	if overrides.Alias != nil {
		resolved.Alias = *overrides.Alias
	}
	if overrides.SystemPrompt != nil {
		resolved.SystemPrompt = *overrides.SystemPrompt
	}
	if overrides.Temperature != nil {
		resolved.Parameters.Temperature = *overrides.Temperature
	}
	if overrides.MaxTokens != nil {
		resolved.Parameters.MaxTokens = *overrides.MaxTokens
	}
	if overrides.Thinking != nil {
		resolved.Parameters.Thinking = *overrides.Thinking
	}
}
