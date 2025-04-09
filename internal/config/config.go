package config

import "os"

func getEnvDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

type Config struct {
	ServerPort   string
	ProxyTimeout int
	TargetURL    string
}

func Load() *Config {
	// TODO Parse args ?
	return &Config{
		ServerPort:   getEnvDefault("PORT", "8080"),
		ProxyTimeout: 30, // seconds
		TargetURL:    getEnvDefault("TARGET_URL", "http://localhost:36000"),
	}
}
