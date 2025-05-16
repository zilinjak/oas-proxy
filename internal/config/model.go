package config

import (
	"net/url"
	"strconv"
)

type Config struct {
	ServerPort   string
	ProxyTimeout int
	TargetURL    *url.URL
	LogLevel     string
	OASPath      string
	StrictMode   bool
}

func (c *Config) String() string {
	return "Config{" +
		"ServerPort: " + c.ServerPort +
		", ProxyTimeout: " + strconv.Itoa(c.ProxyTimeout) +
		", TargetURL: " + c.TargetURL.String() +
		", LogLevel: " + c.LogLevel +
		", OASPath: " + c.OASPath +
		", StrictMode: " + strconv.FormatBool(c.StrictMode) +
		"}"
}
