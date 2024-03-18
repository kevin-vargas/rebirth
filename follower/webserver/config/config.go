package config

import (
	"os"
	"strings"
)

const (
	address = "address"
)

type Config struct {
	Address string
}

func Make() Config {
	defaults := map[string]any{
		address: ":3333",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		Address: defaults[address].(string),
	}
}
