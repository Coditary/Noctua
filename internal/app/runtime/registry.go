package runtime

import (
	"fmt"
	"sort"

	coreerrors "github.com/Coditary/Noctua/internal/core/errors"
	"github.com/Coditary/Noctua/internal/core/ports"
)

type ProviderRegistry struct {
	providers map[string]ports.Provider
}

func NewProviderRegistry() *ProviderRegistry {
	return &ProviderRegistry{providers: map[string]ports.Provider{}}
}

func (r *ProviderRegistry) Register(provider ports.Provider) error {
	if provider == nil {
		return coreerrors.New(coreerrors.KindValidation, "provider is required")
	}

	name := provider.Ref().Name
	if name == "" {
		return coreerrors.New(coreerrors.KindValidation, "provider name is required")
	}

	if _, exists := r.providers[name]; exists {
		return coreerrors.New(coreerrors.KindConflict, fmt.Sprintf("provider %q already registered", name))
	}

	r.providers[name] = provider
	return nil
}

func (r *ProviderRegistry) Get(name string) (ports.Provider, error) {
	provider, ok := r.providers[name]
	if !ok {
		return nil, coreerrors.New(coreerrors.KindNotFound, fmt.Sprintf("provider %q not found", name))
	}

	return provider, nil
}

func (r *ProviderRegistry) Providers() []ports.Provider {
	providers := make([]ports.Provider, 0, len(r.providers))
	for _, provider := range r.providers {
		providers = append(providers, provider)
	}

	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Ref().Name < providers[j].Ref().Name
	})

	return providers
}

func (r *ProviderRegistry) Names() []string {
	providers := r.Providers()
	names := make([]string, 0, len(providers))
	for _, provider := range providers {
		names = append(names, provider.Ref().Name)
	}

	return names
}
