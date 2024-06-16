package config

import (
	"encoding/json"
	"fmt"

	"os"
	"reflect"
	"sync"
)

// ConfigLoader is a generic singleton configuration loader.
type Config[T any] struct {
	once   sync.Once
	config T
	err    error
}

// Load initializes the singleton configuration from a JSON file and environment variables.
func (cl *Config[T]) Load(filePath string) error {
	cl.once.Do(func() {
		cl.config, cl.err = loadConfigFromFileAndEnv[T](filePath)
	})
	return cl.err
}

// Get returns the loaded configuration.
func (cl *Config[T]) Get() T {
	return cl.config
}

// loadConfigFromFileAndEnv reads configuration from a JSON file and environment variables.
func loadConfigFromFileAndEnv[T any](filePath string) (T, error) {
	var config T

	// Read JSON file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		return config, fmt.Errorf("failed to get file: %w", err)
	}

	// Unmarshal JSON data into the config struct
	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Override with environment variables
	overrideWithEnv(&config)

	return config, nil
}

// overrideWithEnv overrides the struct fields with environment variables if they exist.
func overrideWithEnv[T any](config *T) {
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
				// Add other type handling as needed
				// For simplicity, we'll just handle string and int here
			}
		}
	}
}
