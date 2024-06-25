package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const configPath = "%s/config/config-%s.json"

func GetPathEnv() string {
	execPath := getExecutablePath()
	environment := os.Getenv("ENVIRONMENT")
	if len(strings.TrimSpace(environment)) == 0 {
		environment = "local"
	}
	return fmt.Sprintf(configPath, execPath, environment)
}
func GetPathOf(environment string) string {
	execPath := getExecutablePath()
	return fmt.Sprintf(configPath, execPath, environment)
}
func getExecutablePath() string {
	// Get the absolute path of the executable binary
	exePath, err := os.Executable()
	if err != nil {
		panic(err) // Handle error gracefully
	}

	// Get the directory of the executable binary
	exeDir := filepath.ToSlash(filepath.Dir(exePath))

	return exeDir
}
