package runtime

import (
	"testing"

	"github.com/Coditary/Noctua/internal/core/domain"
)

func TestResolveSettingsUsesProfileAndCommandPrecedence(t *testing.T) {
	provider := "remote-stub"
	temperature := 0.8
	maxTokens := 1024
	thinking := true

	config := domain.DefaultConfig()
	config.DefaultProfile = "dev"
	config.Defaults.Provider = "local-stub"
	config.Defaults.Model = "echo-1"
	config.Profiles = map[string]domain.ProfileConfig{
		"dev": {
			Provider:    &provider,
			Temperature: &temperature,
			MaxTokens:   &maxTokens,
		},
	}

	overrideModel := "custom-model"
	settings, err := ResolveSettings(config, domain.CommandOverrides{
		Profile:  "dev",
		Model:    &overrideModel,
		Thinking: &thinking,
	})
	if err != nil {
		t.Fatalf("ResolveSettings() error = %v", err)
	}

	if settings.Provider != "remote-stub" {
		t.Fatalf("provider = %q, want %q", settings.Provider, "remote-stub")
	}
	if settings.Model != "custom-model" {
		t.Fatalf("model = %q, want %q", settings.Model, "custom-model")
	}
	if settings.Parameters.Temperature != 0.8 {
		t.Fatalf("temperature = %v, want 0.8", settings.Parameters.Temperature)
	}
	if settings.Parameters.MaxTokens != 1024 {
		t.Fatalf("max_tokens = %d, want 1024", settings.Parameters.MaxTokens)
	}
	if !settings.Parameters.Thinking {
		t.Fatal("thinking = false, want true")
	}
}

func TestResolveSettingsUsesAliasWhenProviderAndModelAreDefaulted(t *testing.T) {
	alias := "fast"
	config := domain.DefaultConfig()
	config.Defaults.Provider = ""
	config.Defaults.Model = ""
	config.Profiles = map[string]domain.ProfileConfig{
		"default": {
			Alias: &alias,
		},
	}
	config.Aliases = map[string]domain.AliasConfig{
		"fast": {
			Provider: "remote-stub",
			Model:    "gpt-mini",
		},
	}

	settings, err := ResolveSettings(config, domain.CommandOverrides{})
	if err != nil {
		t.Fatalf("ResolveSettings() error = %v", err)
	}

	if settings.Provider != "remote-stub" {
		t.Fatalf("provider = %q, want %q", settings.Provider, "remote-stub")
	}
	if settings.Model != "gpt-mini" {
		t.Fatalf("model = %q, want %q", settings.Model, "gpt-mini")
	}
	if settings.Alias != "fast" {
		t.Fatalf("alias = %q, want %q", settings.Alias, "fast")
	}
}

func TestResolveSettingsRejectsUnknownProfile(t *testing.T) {
	_, err := ResolveSettings(domain.DefaultConfig(), domain.CommandOverrides{Profile: "missing"})
	if err == nil {
		t.Fatal("ResolveSettings() error = nil, want error")
	}
}
