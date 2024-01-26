package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/kevin-vargas/rebirth/dns"
)

const (
	update_tag = "rebirth-ip"
)

type dnsI struct {
	token string
}

func (d *dnsI) Update(ctx context.Context, ip string) error {
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
			if record.Comment == update_tag {
				_, err := api.UpdateDNSRecord(ctx, zoneID, cloudflare.UpdateDNSRecordParams{
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
