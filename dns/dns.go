package dns

import "context"

type DNSOps struct {
	Proxied bool
}

type DNS interface {
	Update(context.Context, string, DNSOps) error
}
