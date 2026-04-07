package stub

import (
	"context"
	"fmt"

	"github.com/Coditary/Noctua/internal/core/domain"
)

type Provider struct {
	ref          domain.ProviderRef
	capabilities domain.CapabilitySet
	models       []string
}

func NewLocal() *Provider {
	return &Provider{
		ref: domain.ProviderRef{Name: "local-stub"},
		capabilities: domain.CapabilitySet{
			Chat:      true,
			Streaming: false,
			Local:     true,
		},
		models: []string{"echo-1", "reason-1"},
	}
}

func NewRemote() *Provider {
	return &Provider{
		ref: domain.ProviderRef{Name: "remote-stub"},
		capabilities: domain.CapabilitySet{
			Chat:      true,
			Streaming: true,
			Remote:    true,
		},
		models: []string{"gpt-mini", "vision-mini"},
	}
}

func (p *Provider) Ref() domain.ProviderRef {
	return p.ref
}

func (p *Provider) Capabilities() domain.CapabilitySet {
	return p.capabilities
}

func (p *Provider) Execute(_ context.Context, request domain.ExecutionRequest) (domain.ExecutionResult, error) {
	prompt := ""
	if len(request.Messages) > 0 {
		prompt = domain.MessageText(request.Messages[len(request.Messages)-1])
	}

	response := fmt.Sprintf("[%s/%s] %s", p.ref.Name, request.Model.Name, prompt)

	return domain.ExecutionResult{
		Provider: p.ref,
		Model:    request.Model,
		Message:  domain.AssistantMessage(response),
		Usage: domain.Usage{
			InputTokens:  len(prompt),
			OutputTokens: len(response),
		},
	}, nil
}

func (p *Provider) ListModels(context.Context) ([]domain.ModelRef, error) {
	models := make([]domain.ModelRef, 0, len(p.models))
	for _, model := range p.models {
		models = append(models, domain.ModelRef{Provider: p.ref.Name, Name: model})
	}

	return models, nil
}
