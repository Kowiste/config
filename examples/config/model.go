package config

type MyConfig struct {
	DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
	Port        int    `json:"port" env:"PORT"`
}
type MyConfig2 struct {
	DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
	Port        string `json:"port" env:"PORT"`
}