package app

import "flag"

type (
	Config struct {
		Server
		Logging
	}

	Server struct {
		JSONAPIPort int
	}

	Logging struct {
		Level string
	}
)

func LoadConfig() *Config {
	c := new(Config)

	// Server
	flag.IntVar(&c.Server.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	// Logging
	flag.StringVar(&c.Logging.Level, "logging-level", "error", "Logging level")

	return c
}
