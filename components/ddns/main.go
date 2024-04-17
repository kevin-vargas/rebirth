package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevin-vargas/rebirth/components/ddns/config"
	"github.com/kevin-vargas/rebirth/dns"
	"github.com/kevin-vargas/rebirth/dns/cloudflare"
	"github.com/kevin-vargas/rebirth/ip"
	"github.com/kevin-vargas/rebirth/ip/public"
)

func main() {
	godotenv.Load()
	cfg := config.Make()
	d := cloudflare.New(cfg.CloudflareAPIToken)
	i := public.New()
	ctx := context.Background()
	setIP := makeSetIP()
	for range time.Tick(cfg.Period) {
		_, err := setIP(ctx, cfg, i, d)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Done")
}

func makeSetIP() func(context.Context, config.Config, ip.IP, dns.DNS) (string, error) {
	var currentIP string
	return func(ctx context.Context, cfg config.Config, i ip.IP, d dns.DNS) (string, error) {
		ip, err := i.Get()
		if err != nil {
			return "", err
		}
		if currentIP != "" && ip == currentIP {
			return "", nil
		}
		currentIP = ip
		if err := d.UpdateSingle(ctx, ip, cfg.Entry, cfg.EntryZone, dns.DNSOps{Proxied: false}); err != nil {
			return "", err
		}
		return ip, nil
	}
}
