package config

type MyConfig struct {
	DatabaseURL  string `json:"DatabaseURL" env:"DATABASE_URL"`
	DatabasePort int    `json:"DatabasePort" env:"DATABASE_PORT"`
	ServiceName  string `json:"ServiceName" env:"SERVICE_URL"`
	ServicePort  int    `json:"ServicePort" env:"SERVICE_PORT"`
}
type MyConfig2 struct {
	DatabaseURL string `json:"database_url" env:"DATABASE_URL"`
	Port        string `json:"port" env:"PORT"`
}
