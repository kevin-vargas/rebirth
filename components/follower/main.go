package main

import (
	"context"
	"errors"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevin-vargas/rebirth/alive"
	"github.com/kevin-vargas/rebirth/alive/ping"
	"github.com/kevin-vargas/rebirth/components/follower/config"
	"github.com/kevin-vargas/rebirth/db"
	"github.com/kevin-vargas/rebirth/db/redis"
	"github.com/kevin-vargas/rebirth/dns"
	"github.com/kevin-vargas/rebirth/dns/cloudflare"
	"github.com/kevin-vargas/rebirth/ip"
	"github.com/kevin-vargas/rebirth/ip/public"
)

type manager struct {
	db db.DB
	a  alive.Aliver
	i  ip.IP
	d  dns.DNS
}

func (m *manager) needToTakeControll(ctx context.Context) (b bool, err error) {
	// TODO: make sense that the master can not be empty but the current ip can?
	currentIP, err := m.db.GetCurrentIP(ctx)
	if err != nil && !errors.Is(err, db.ErrNotFound) {
		return
	}
	publicIP, err := m.i.Get()
	if err != nil {
		return
	}
	if currentIP == publicIP {
		return false, nil
	}
	ipMaster, err := m.db.GetMasterIP(ctx)
	if errors.Is(err, db.ErrNotFound) {
		return true, nil
	}
	if err != nil {
		return
	}
	isAlive, err := m.a.IsAlive(ipMaster)
	if err != nil {
		return
	}
	if isAlive {
		return false, nil
	}
	// TODO: check concurrent check context cancel get public ip

	return true, nil
}

func (m *manager) takeControll(ctx context.Context) error {
	ops := dns.DNSOps{
		Proxied: true,
	}
	ip, err := m.i.Get()
	if err != nil {
		return err
	}
	if err := m.d.Update(ctx, ip, ops); err != nil {
		return err
	}
	if err := m.db.SetCurrentIP(ctx, ip); err != nil {
		return err
	}
	return nil
}

const (
	count        = 3
	keyMaster    = "rebirth-master"
	keyCurrentIP = "rebirth-current-ip"
)

func main() {
	godotenv.Load()
	cfg := config.Make()

	ctx := context.Background()
	d := cloudflare.New(cfg.CloudflareAPIToken)
	db := redis.New(
		cfg.RedisAddress,
		cfg.RedisPassword,
		keyMaster,
		keyCurrentIP,
	)
	a, err := ping.New(count)
	if err != nil {
		panic(err)
	}
	i := public.New()
	m := &manager{
		db: db,
		a:  a,
		i:  i,
		d:  d,
	}
	for range time.Tick(cfg.Period) {
		n, err := m.needToTakeControll(ctx)
		if err != nil {
			panic(err)
		}
		if n {
			err := m.takeControll(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
