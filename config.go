package main

import (
	"encoding/json"
	"fmt"

	"os"
	"reflect"
	"sync"
)

// ConfigLoader is a generic singleton configuration loader.
type config[T any] struct {
	once   sync.Once
	config T
}

// New creates a new instance of Config for a path.
func New[T any](filePath string) (cl config[T], err error) {
	cl.once.Do(func() {
		cl.config, err = fromFile[T](filePath)
	})
	return
}

// Get returns the loaded configuration.
func (cl *config[T]) Get() T {
	return cl.config
}

// loadConfigFromFileAndEnv reads configuration from a JSON file and environment variables.
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


