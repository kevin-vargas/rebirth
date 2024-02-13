package config

import (
	"os"
	"strings"
)

const (
	cloudflare_api_token = "cloudflare_api_token"
	redis_address        = "redis_address"
	redis_password       = "redis_password"
)

type Config struct {
	RedisAddress       string
	RedisPassword      string
	CloudflareAPIToken string
}

func Make() Config {
	defaults := map[string]any{
		redis_address:        "localhost",
		redis_password:       "",
		cloudflare_api_token: "",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		RedisAddress:       defaults[redis_address].(string),
		RedisPassword:      defaults[redis_password].(string),
		CloudflareAPIToken: defaults[cloudflare_api_token].(string),
	}
}
