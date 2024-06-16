package main

import (
	"fmt"

	"github.com/kowiste/config"
)

type MyConfig struct {
	DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
	Port        int    `json:"port" env:"PORT"`
}
type MyConfig2 struct {
	DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
	Port        string `json:"port" env:"PORT"`
}

func main() {
	config, err := config.New[MyConfig2]("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	config.Get()
}