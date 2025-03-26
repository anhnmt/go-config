package defaultconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultConfigType = "yml"
	defaultConfigName = "dev"
	defaultConfigPath = "."
)

func init() {
	viper.SetConfigType(defaultConfigType)
	viper.SetConfigName(defaultConfigName)
	viper.AddConfigPath(defaultConfigPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// Load loads the config file and binds values to cfg.
//
// Example usage:
//
//	if err := config.Load(dir, env, &cfg) err != nil {
//	    ...
//	}
func Load(dir, env string, cfg any) error {
	v := reflect.ValueOf(cfg)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("cfg must be a non-nil pointer to a struct, got %T", cfg)
	}

	if dir != "" && dir != defaultConfigPath {
		viper.AddConfigPath(dir)
	}

	if env != "" && dir != defaultConfigName {
		viper.SetConfigName(env)
	}

	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return fmt.Errorf("fatal error reading config file: %w", err)
	}

	bindValues(cfg)

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	return nil
}
