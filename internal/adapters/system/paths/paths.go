package paths

import (
	"os"
	"path/filepath"

	"github.com/Coditary/Noctua/internal/core/domain"
	coreerrors "github.com/Coditary/Noctua/internal/core/errors"
)

type Directories struct {
	ConfigDir string `json:"config_dir"`
	DataDir   string `json:"data_dir"`
	CacheDir  string `json:"cache_dir"`
	LogDir    string `json:"log_dir"`
}

type Resolver struct {
	appName string
}

func NewResolver(appName string) Resolver {
	return Resolver{appName: appName}
}

func (r Resolver) Resolve(config domain.PathsConfig) (Directories, error) {
	configBase, err := os.UserConfigDir()
	if err != nil {
		return Directories{}, coreerrors.Wrap(coreerrors.KindConfiguration, "could not determine user config directory", err)
	}

	cacheBase, err := os.UserCacheDir()
	if err != nil {
		return Directories{}, coreerrors.Wrap(coreerrors.KindConfiguration, "could not determine user cache directory", err)
	}

	dataBase := os.Getenv("XDG_DATA_HOME")
	if dataBase == "" {
		home, homeErr := os.UserHomeDir()
		if homeErr != nil {
			return Directories{}, coreerrors.Wrap(coreerrors.KindConfiguration, "could not determine user home directory", homeErr)
		}

		dataBase = filepath.Join(home, ".local", "share")
	}

	dirs := Directories{
		ConfigDir: firstNonEmpty(os.Getenv("NOCTUA_CONFIG_DIR"), config.ConfigDir, filepath.Join(configBase, r.appName)),
		DataDir:   firstNonEmpty(os.Getenv("NOCTUA_DATA_DIR"), config.DataDir, filepath.Join(dataBase, r.appName)),
		CacheDir:  firstNonEmpty(os.Getenv("NOCTUA_CACHE_DIR"), config.CacheDir, filepath.Join(cacheBase, r.appName)),
	}
	dirs.LogDir = firstNonEmpty(os.Getenv("NOCTUA_LOG_DIR"), config.LogDir, filepath.Join(dirs.DataDir, "logs"))

	return dirs, nil
}

func (d Directories) Bootstrap() error {
	for _, path := range []string{d.ConfigDir, d.DataDir, d.CacheDir, d.LogDir} {
		if path == "" {
			continue
		}

		if err := os.MkdirAll(path, 0o755); err != nil {
			return coreerrors.Wrap(coreerrors.KindConfiguration, "could not create application directories", err)
		}
	}

	return nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}

	return ""
}
