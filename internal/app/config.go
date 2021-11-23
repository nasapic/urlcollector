package app

import (
	"flag"
	"os"
	"strconv"
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
		APIKEYEnvar        string
		APIKEY             string
		MaxConcurrentEnvar string
		MaxConcurrent      int
		TimeoutInSecs      int
	}
)

const (
	nasaDefaultAPIKEY      = "DEMO_KEY"
	nasaDefaultMaxRequests = 5
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
	flag.StringVar(&cfg.NASAAPI.MaxConcurrentEnvar, "nasa-api-max-concurrent-envar", "CONCURRENT_REQUESTS", "NASA API max concurrent requests envar")
	flag.IntVar(&cfg.NASAAPI.TimeoutInSecs, "nasa-api-timeout-in-secs", 5, "NASA API timeout in secs")

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

	maxReq := toInt(loadEnvar(na.MaxConcurrentEnvar))
	if maxReq < 1 {
		maxReq = nasaDefaultMaxRequests
	}

	na.MaxConcurrent = maxReq

}

func loadEnvar(name string) string {
	return os.Getenv(name)
}

func toInt(intString string) int {
	i, err := strconv.Atoi(intString)
	if err == nil {
		return 0
	}

	return i
}
