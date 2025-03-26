package defaultconfig

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func bindValues(iface any) {
	v := reflect.Indirect(reflect.ValueOf(iface))
	processField(v, v.Type())
}

func processField(v reflect.Value, t reflect.Type, parts ...string) {
	partsLen := len(parts)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		tag, hasTag := field.Tag.Lookup("mapstructure")
		if !hasTag {
			continue
		}

		key := tag
		if partsLen > 0 {
			key = strings.Join(append(parts, tag), ".")
		}
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

		if err := viper.BindEnv(key, envKey); err != nil {
			fmt.Printf("Warning: Failed to bind environment variable %s: %v\n", envKey, err)
		}

		if !viper.IsSet(key) {
			if value, hasDefault := field.Tag.Lookup("default"); hasDefault {
				viper.SetDefault(key, value)
			}
		}

		if fieldVal.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(field.Type.Elem()))
			}
			fieldVal = fieldVal.Elem()
		}

		if fieldVal.Kind() == reflect.Struct {
			processField(fieldVal, fieldVal.Type(), append(parts, tag)...)
		}
	}
}
