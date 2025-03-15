package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var File = "dev.yml"

// InitConfig loads the config file and binds values to cfg.
//
// Example usage:
//
//	if err := config.InitConfig(dir, env, &cfg) err != nil {
//	    ...
//	}
func InitConfig(dir, env string, cfg any) error {
	v := reflect.ValueOf(cfg)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("cfg must be a non-nil pointer to a struct, got %T", cfg)
	}

	path, err := filePath(dir, env)
	if err != nil {
		return fmt.Errorf("get config file path error: %w", err)
	}

	viper.SetConfigFile(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configNotFound) {
			return fmt.Errorf("fatal error reading config file: %w", err)
		}
	}

	bindValues(cfg)

	if err = viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	return nil
}

func filePath(dir, env string) (string, error) {
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("getwd error: %w", err)
		}
	}

	if env != "" {
		File = env + ".yml"
	}

	return filepath.ToSlash(filepath.Join(dir, File)), nil
}
