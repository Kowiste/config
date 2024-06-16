package config

import "fmt"

const configPath = "./config/config-%s"

func GetPathEnv(environment string) string {
	return fmt.Sprintf(configPath, environment)
}
