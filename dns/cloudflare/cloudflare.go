package cloudflare

import (
	"context"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/kevin-vargas/rebirth/dns"
)

const (
	update_tag             = "rebirth-ip"
	update_tag_proxied     = "rebirth-ip-proxied"
	update_tag_non_proxied = "rebirth-ip-non-proxied"
)

type dnsI struct {
	token string
}

func (d *dnsI) UpdateSingle(ctx context.Context, ip, url, zone string, ops dns.DNSOps) error {
	api, err := cloudflare.NewWithAPIToken(d.token)
	if err != nil {
		return err
	}
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		return err
	}
	cZoneID := cloudflare.ZoneIdentifier(zoneID)
	r, _, err := api.ListDNSRecords(
		ctx,
		cZoneID,
		cloudflare.ListDNSRecordsParams{
			Name: url,
		},
	)
	if err != nil {
		return err
	}
	for _, record := range r {
		_, err := api.UpdateDNSRecord(ctx, cZoneID, cloudflare.UpdateDNSRecordParams{
			Proxied: &ops.Proxied,
			ID:      record.ID,
			Content: ip,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *dnsI) Update(ctx context.Context, ip string, ops dns.DNSOps) error {
	api, err := cloudflare.NewWithAPIToken(d.token)
	if err != nil {
		return err
	}
	z, err := api.ListZones(ctx)
	if err != nil {
		return err
	}
	for _, zone := range z {
		// TODO: use page res
		zoneID := cloudflare.ZoneIdentifier(zone.ID)
		r, _, err := api.ListDNSRecords(
			ctx,
			zoneID,
			cloudflare.ListDNSRecordsParams{},
		)
		if err != nil {
			return err
		}
		for _, record := range r {
			if strings.Contains(record.Comment, update_tag) {
				isProxied := ops.Proxied
				switch record.Comment {
				case update_tag_non_proxied:
					isProxied = false
				case update_tag_proxied:
					isProxied = true
				}
				_, err := api.UpdateDNSRecord(ctx, zoneID, cloudflare.UpdateDNSRecordParams{
					Proxied: &isProxied,
					ID:      record.ID,
					Content: ip,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func New(t string) dns.DNS {
	return &dnsI{token: t}
}
