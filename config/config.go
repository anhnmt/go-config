package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var File = "dev.yml"

// InitConfig load config file and bind values to cfg
//
// err := config.InitConfig(dir, env, &cfg)
//
//	if err != nil {
//	    ...
//	}
func InitConfig(dir, env string, cfg any) error {
	path, err := filePath(dir, env)
	if err != nil {
		return fmt.Errorf("get config file path error: %w", err)
	}

	viper.SetConfigFile(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var configFileNotFoundError viper.ConfigFileNotFoundError
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil && !errors.As(err, &configFileNotFoundError) {
		return fmt.Errorf("fatal error config file: %w", err)
	}

	bindValues(cfg)

	err = viper.Unmarshal(cfg)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
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
