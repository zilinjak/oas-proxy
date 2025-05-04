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
	targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://tom.preston-werner.com"))
	//targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://preview-k3w-it-14824411752.t.all.web.ne.kosik.systems"))
	//targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "https://httpbin.zilinek.fun"))
	//targetUrl, err := url.Parse(getEnvDefault("TARGET_URL", "http://localhost:8000"))

	if err != nil {
		panic("Invalid TARGET_URL: " + err.Error())
	}

	AppConfig = &Config{
		ServerPort:   getEnvDefault("PORT", "8080"),
		ProxyTimeout: 30,
		TargetURL:    targetUrl,
		LogLevel:     getEnvDefault("LOG_LEVEL", "DEBUG"),
	}

	return AppConfig
}
