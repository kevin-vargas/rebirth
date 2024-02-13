package public

import (
	"errors"
	"io"
	"net/http"

	"github.com/kevin-vargas/rebirth/ip"
)

const (
	ip_service = "https://api.ipify.org"
)

type ipPublic struct{}

func (i *ipPublic) Get() (string, error) {
	res, err := http.Get(ip_service)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("invalid status code on ip service")
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func New() ip.IP {
	return &ipPublic{}
}
