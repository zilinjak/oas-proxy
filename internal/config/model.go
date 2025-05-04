package config

import "net/url"

type Config struct {
	ServerPort   string
	ProxyTimeout int
	TargetURL    *url.URL
	LogLevel     string
}

func (c *Config) String() string {
	return "Config{" +
		"ServerPort: " + c.ServerPort +
		", ProxyTimeout: " + string(c.ProxyTimeout) +
		", TargetURL: " + c.TargetURL.String() +
		", LogLevel: " + c.LogLevel +
		"}"
}
