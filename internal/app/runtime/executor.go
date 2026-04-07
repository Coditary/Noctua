package runtime

import (
	"context"

	"github.com/Coditary/Noctua/internal/core/domain"
)

type RunResult struct {
	Settings domain.ResolvedSettings `json:"settings"`
	Result   domain.ExecutionResult  `json:"result"`
}

type Executor struct {
	registry *ProviderRegistry
}

func NewExecutor(registry *ProviderRegistry) *Executor {
	return &Executor{registry: registry}
}

func (e *Executor) RunPrompt(ctx context.Context, config domain.AppConfig, overrides domain.CommandOverrides, prompt string) (RunResult, error) {
	settings, err := ResolveSettings(config, overrides)
	if err != nil {
		return RunResult{}, err
	}

	provider, err := e.registry.Get(settings.Provider)
	if err != nil {
		return RunResult{}, err
	}

	request := domain.ExecutionRequest{
		Model: domain.ModelRef{
			Provider: settings.Provider,
			Name:     settings.Model,
			Alias:    settings.Alias,
		},
		SystemPrompt: settings.SystemPrompt,
		Parameters:   settings.Parameters,
		Messages:     []domain.Message{domain.UserMessage(prompt)},
	}

	result, err := provider.Execute(ctx, request)
	if err != nil {
		return RunResult{}, err
	}

	return RunResult{Settings: settings, Result: result}, nil
}

func (e *Executor) ListModels(ctx context.Context) ([]domain.ModelRef, error) {
	models := make([]domain.ModelRef, 0)
	for _, provider := range e.registry.Providers() {
		catalog, ok := provider.(interface {
			ListModels(context.Context) ([]domain.ModelRef, error)
		})
		if !ok {
			continue
		}

		listed, err := catalog.ListModels(ctx)
		if err != nil {
			return nil, err
		}

		models = append(models, listed...)
	}

	return models, nil
}
