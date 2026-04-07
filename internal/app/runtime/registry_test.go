package runtime

import (
	"testing"

	"github.com/Coditary/Noctua/internal/adapters/providers/stub"
)

func TestProviderRegistryRejectsDuplicateProviders(t *testing.T) {
	registry := NewProviderRegistry()
	if err := registry.Register(stub.NewLocal()); err != nil {
		t.Fatalf("Register() first error = %v", err)
	}

	if err := registry.Register(stub.NewLocal()); err == nil {
		t.Fatal("Register() second error = nil, want duplicate error")
	}
}

func TestProviderRegistryListsSortedNames(t *testing.T) {
	registry := NewProviderRegistry()
	_ = registry.Register(stub.NewRemote())
	_ = registry.Register(stub.NewLocal())

	names := registry.Names()
	if len(names) != 2 {
		t.Fatalf("len(names) = %d, want 2", len(names))
	}
	if names[0] != "local-stub" || names[1] != "remote-stub" {
		t.Fatalf("names = %v, want sorted provider names", names)
	}
}
