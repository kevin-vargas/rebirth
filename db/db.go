package db

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

type DB interface {
	GetMasterIP(context.Context) (string, error)
	GetCurrentIP(context.Context) (string, error)
	SetCurrentIP(context.Context, string) error
}
