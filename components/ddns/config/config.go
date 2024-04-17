package config

import (
	"os"
	"strings"
	"time"
)

const (
	entry                = "entry"
	entry_zone           = "entry_zone"
	period               = "period"
	cloudflare_api_token = "cloudflare_api_token"
)

type Config struct {
	EntryZone          string
	Entry              string
	Period             time.Duration
	CloudflareAPIToken string
}

func Make() Config {
	defaults := map[string]any{
		entry:                "",
		entry_zone:           "",
		period:               5 * 60 * 1000, // 5 mins
		cloudflare_api_token: "",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		EntryZone:          defaults[entry_zone].(string),
		Entry:              defaults[entry].(string),
		Period:             time.Duration(defaults[period].(int)) * time.Millisecond,
		CloudflareAPIToken: defaults[cloudflare_api_token].(string),
	}
}
