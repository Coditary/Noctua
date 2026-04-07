package stub

import (
	"context"
	"testing"

	"github.com/Coditary/Noctua/internal/core/domain"
)

func TestLocalProviderExecutesPrompt(t *testing.T) {
	provider := NewLocal()
	result, err := provider.Execute(context.Background(), domain.ExecutionRequest{
		Model:    domain.ModelRef{Provider: provider.Ref().Name, Name: "echo-1"},
		Messages: []domain.Message{domain.UserMessage("hello")},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if result.Provider.Name != "local-stub" {
		t.Fatalf("provider = %q, want %q", result.Provider.Name, "local-stub")
	}
	if text := domain.MessageText(result.Message); text == "" {
		t.Fatal("response text is empty")
	}
}

func TestRemoteProviderListsModels(t *testing.T) {
	provider := NewRemote()
	models, err := provider.ListModels(context.Background())
	if err != nil {
		t.Fatalf("ListModels() error = %v", err)
	}

	if len(models) == 0 {
		t.Fatal("ListModels() returned no models")
	}
	if !provider.Capabilities().Remote {
		t.Fatal("remote capability = false, want true")
	}
}
