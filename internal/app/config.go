package app

import (
	"flag"
	"os"
	"strings"
)

type (
	Config struct {
		Server
		Logging
		NASAAPI
	}

	Server struct {
		JSONAPIPort int
	}

	Logging struct {
		Level string
	}

	NASAAPI struct {
		APIKEYEnvar string
		APIKEY      string
	}
)

const (
	nasaDefaultAPIKEY = "DEMO_KEY"
)

// NOTE: A little update eventually would allow all envar names to be passed
// as parameters instead of the actual values.
func LoadConfig() *Config {
	cfg := Config{}

	// Server
	flag.IntVar(&cfg.Server.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	// Logging
	flag.StringVar(&cfg.Logging.Level, "logging-level", "error", "Logging level")

	// NASA API
	flag.StringVar(&cfg.NASAAPI.APIKEYEnvar, "nasa-api-key-envar", "", "NASA API Key envar")

	flag.Parse()

	cfg.LoadFromEnvar()

	return &cfg
}

func (cfg *Config) LoadFromEnvar() {
	cfg.NASAAPI.LoadFromEnvar()
}

func (na *NASAAPI) LoadFromEnvar() {
	ak := loadEnvar(na.APIKEYEnvar)

	if "" == strings.TrimSpace(ak) {
		ak = nasaDefaultAPIKEY
	}

	na.APIKEY = ak
}

func loadEnvar(name string) string {
	return os.Getenv(name)
}
