package config

import (
	"os"
	"strings"
)

const (
	cloudflare_api_token = "cloudflare_api_token"
)

type Config struct {
	CloudflareAPIToken string
}

func Make() Config {
	defaults := map[string]any{
		cloudflare_api_token: "",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		CloudflareAPIToken: defaults[cloudflare_api_token].(string),
	}
}
