package config

import (
	"os"
	"strings"
	"time"
)

const (
	cloudflare_api_token = "cloudflare_api_token"
	redis_address        = "redis_address"
	redis_password       = "redis_password"
	period               = "period"
)

type Config struct {
	Address            string
	RedisAddress       string
	RedisPassword      string
	CloudflareAPIToken string
	Period             time.Duration
}

func Make() Config {
	defaults := map[string]any{
		redis_address:        "localhost",
		redis_password:       "",
		cloudflare_api_token: "",
		period:               5 * 60 * 1000, // 5 mins
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
		Period:             time.Duration(defaults[period].(int)) * time.Millisecond,
	}
}
