// Package path implements utility to read configuration
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync"
)

// Config is a singleton configuration loader.
type Config[T any] struct {
	once   sync.Once
	config T
}

var instance *Config[any]
var mu sync.Mutex

// New initializes the configuration loader for a given file path.
// It will only load the configuration once.
func New[T any](filePath string) (err error) {
	mu.Lock()
	defer mu.Unlock()

	// Check if an instance already exists
	if instance != nil {
		instance.config, err = fromFile[T](filePath)
		return
	}

	// Create a new instance
	instance = &Config[any]{}
	instance.once.Do(func() {
		instance.config, err = fromFile[T](filePath)
	})

	return
}

// Get returns the loaded configuration.
func Get[T any]() (config T, err error) {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		return config, fmt.Errorf("configuration not loaded")
	}

	// Type assert the stored config to the desired type
	config, ok := instance.config.(T)
	if !ok {
		return config, fmt.Errorf("configuration type mismatch")
	}

	return config, err
}

// fromFile reads configuration from a JSON file and environment variables.
func fromFile[T any](filePath string) (T, error) {
	var config T

	// Read JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to get file: %w", err)
	}

	// Unmarshal JSON data into the config struct
	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Override with environment variables
	overrideEnv(&config)

	return config, nil
}

// overrideWithEnv overrides the struct fields with environment variables if they exist.
func overrideEnv[T any](config *T) {
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		envVar := fieldType.Tag.Get("env")
		if envVar == "" {
			continue
		}

		if envValue, exists := os.LookupEnv(envVar); exists {
			switch field.Kind() {
			case reflect.String:
				field.SetString(envValue)
			case reflect.Int:
				if intValue, err := strconv.Atoi(envValue); err == nil {
					field.SetInt(int64(intValue))
				} else {
					// Handle error (e.g., log it or return an error)
				}
			case reflect.Bool:
				if boolValue, err := strconv.ParseBool(envValue); err == nil {
					field.SetBool(boolValue)
				} else {
					// Handle error (e.g., log it or return an error)
				}
			case reflect.Float32, reflect.Float64:
				if floatValue, err := strconv.ParseFloat(envValue, field.Type().Bits()); err == nil {
					field.SetFloat(floatValue)
				} else {
					// Handle error (e.g., log it or return an error)
				}
			}
		}
	}
}
