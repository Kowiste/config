package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configPath = "%s/config/config-%s"

func GetPathEnv(environment string) string {
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
