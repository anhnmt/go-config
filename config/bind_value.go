package config

import (
    "reflect"
    "strings"

    "github.com/spf13/viper"
)

func unwrapPointer(iface any) (reflect.Value, reflect.Type) {
    v := reflect.ValueOf(iface)
    t := reflect.TypeOf(iface)

    for v.Kind() == reflect.Pointer {
        v = v.Elem()
        t = t.Elem()
    }

    return v, t
}

func bindValues(iface any, parts ...string) {
    ifv, ift := unwrapPointer(iface)
    processField(ifv, ift, parts)
}

func processField(v reflect.Value, t reflect.Type, parts []string) {
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldVal := v.Field(i)

        tag, ok := field.Tag.Lookup("mapstructure")
        if !ok {
            continue
        }

        var builder strings.Builder
        if len(parts) > 0 {
            builder.WriteString(strings.Join(parts, "."))
            builder.WriteString(".")
        }
        builder.WriteString(tag)

        key := builder.String()
        envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))

        _ = viper.BindEnv(key, envKey)

        if value, hasDefault := field.Tag.Lookup("defaultvalue"); hasDefault {
            viper.SetDefault(key, value)
        }

        if fieldVal.Kind() == reflect.Struct {
            processField(fieldVal, field.Type, append(parts, tag))
        }
    }
}
