package config

import (
	"net/url"
	"os"
)

var AppConfig *Config

func init() {
	AppConfig = load()
}

func getEnvDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func load() *Config {
	// TODO Parse args ?
	// TODO use https://github.com/knadh/koanf
	//targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://test.k6.io"))
	// targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://.t.all.web.ne.kosik.systems"))
	//targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://httpbin.zilinek.fun"))
	targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "http://localhost:8000"))

	if err != nil {
		panic("Invalid TARGET_URL: " + err.Error())
	}

	AppConfig = &Config{
		ServerPort:   getEnvDefault("PORT", "8080"),
		ProxyTimeout: 30,
		TargetURL:    targetUrl,
		LogLevel:     getEnvDefault("LOG_LEVEL", "DEBUG"),
		OASPath:      getEnvDefault("OAS_PATH", "./tests/oas/openapi.yaml"),
	}

	return AppConfig
}
