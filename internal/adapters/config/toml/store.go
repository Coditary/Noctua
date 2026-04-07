package toml

import (
	"context"
	"os"

	toml "github.com/BurntSushi/toml"

	"github.com/Coditary/Noctua/internal/core/domain"
	coreerrors "github.com/Coditary/Noctua/internal/core/errors"
)

type Store struct {
	path string
}

func New(path string) Store {
	return Store{path: path}
}

func (s Store) Path() string {
	return s.path
}

func (s Store) Load(context.Context) (domain.AppConfig, error) {
	config, err := s.LoadOptional(context.Background())
	if err != nil {
		return domain.AppConfig{}, err
	}

	return config, nil
}

func (s Store) LoadOptional(context.Context) (domain.AppConfig, error) {
	config := domain.DefaultConfig()
	if s.path == "" {
		return config, nil
	}

	if _, err := os.Stat(s.path); err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}

		return domain.AppConfig{}, coreerrors.Wrap(coreerrors.KindConfiguration, "could not stat config file", err)
	}

	if _, err := toml.DecodeFile(s.path, &config); err != nil {
		return domain.AppConfig{}, coreerrors.Wrap(coreerrors.KindConfiguration, "could not decode TOML config", err)
	}

	if config.Profiles == nil {
		config.Profiles = map[string]domain.ProfileConfig{}
	}
	if config.Aliases == nil {
		config.Aliases = map[string]domain.AliasConfig{}
	}

	return config, nil
}
