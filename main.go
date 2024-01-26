package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/kevin-vargas/rebirth/config"
	"github.com/kevin-vargas/rebirth/dns/cloudflare"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	cfg := config.Make()
	dns := cloudflare.New(cfg.CloudflareAPIToken)
	err := dns.Update(ctx, "192.168.0.1")
	if err != nil {
		log.Fatal(err)
	}
}
