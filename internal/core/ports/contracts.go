package ports

import (
	"context"

	"github.com/Coditary/Noctua/internal/core/domain"
)

type Provider interface {
	Ref() domain.ProviderRef
	Capabilities() domain.CapabilitySet
	Execute(context.Context, domain.ExecutionRequest) (domain.ExecutionResult, error)
}

type CatalogProvider interface {
	ListModels(context.Context) ([]domain.ModelRef, error)
}

type ConfigStore interface {
	Load(context.Context) (domain.AppConfig, error)
}

type SessionStore interface {
	Save(context.Context, domain.Session) error
	Load(context.Context, string) (domain.Session, error)
}

type ResourceProbe interface {
	Snapshot(context.Context) (domain.ResourceSnapshot, error)
}
