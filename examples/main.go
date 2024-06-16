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
	err := config.New[MyConfig2](config.G)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	cfg, err := config.Get[MyConfig2]()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Print(cfg)
}
