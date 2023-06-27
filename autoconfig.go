package shims

import (
	"fmt"
	"reflect"
)

// DefaultValue determines if the input is value (not empty), and if it is,
// returns the input value, otherwise returns a provided default.
func DefaultValue[T comparable](in, def T) T {
	if reflect.ValueOf(in).IsZero() {
		return def
	}
	return in
}

func AutoConfigure(src any, dst any) error {
	data := autoconfigread(src)
	return autoconfigwrite(dst, data)
}

// autoconfigread reads values from a structure marked with autoconfig tags
func autoconfigread(v any) map[string]any {
	out := map[string]any{}

	var val reflect.Value
	if reflect.TypeOf(v).Name() != "" {
		val = reflect.ValueOf(v)
	} else {
		val = reflect.ValueOf(v).Elem()
	}

	fmt.Printf("Found %d fields\n", val.NumField())
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		if typeField.Type.Kind() == reflect.Struct {
			// Recurse into child structure
			sub := autoconfigread(val.Field(i))
			// Combine into output structure
			for ksub, vsub := range sub {
				out[ksub] = vsub
			}
			// Skip anything else we might do
			continue
		}

		// Value fields, pull autoconfig tag
		n := tag.Get("autoconfig")
		if len(n) < 1 {
			fmt.Println("Bailing out, invalid autoconfig tag ", n)
			continue
		}

		out[n] = val.Field(i)
	}

	return out
}

func autoconfigwrite(v any, data map[string]any) error {
	var val reflect.Value
	if reflect.TypeOf(v).Name() != "" {
		val = reflect.ValueOf(v)
	} else {
		val = reflect.ValueOf(v).Elem()
	}

	fmt.Printf("Found %d fields\n", val.NumField())
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		if typeField.Type.Kind() == reflect.Struct {
			// Recurse into structure with data
			err := autoconfigwrite(val.Field(i), data)
			// Report any errors
			if err != nil {
				fmt.Printf("ERR: %s\n", err.Error())
			}
			// Skip value fields processing
			continue
		}

		n := tag.Get("autoconfig")
		if len(n) < 1 {
			fmt.Println("Bailing out, invalid autoconfig tag ", n)
			continue
		}

		if !val.Field(i).CanSet() {
			fmt.Printf("Can't set field %s\n", val.Field(i).String())
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.String:
			val.Field(i).SetString(data[n].(string))
		case reflect.Bool:
			val.Field(i).SetBool(data[n].(bool))
		case reflect.Int:
			val.Field(i).SetInt(data[n].(int64))
		default:
		}
	}

	return nil
}
