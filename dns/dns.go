package dns

import "context"

type DNSOps struct {
	Proxied bool
}

type DNS interface {
	UpdateSingle(context.Context, string, string, string, DNSOps) error
	Update(context.Context, string, DNSOps) error
}
