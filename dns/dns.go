package dns

import "context"

type DNS interface {
	Update(context.Context, string) error
}
