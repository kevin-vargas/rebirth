package ping

import (
	goping "github.com/go-ping/ping"
	"github.com/kevin-vargas/rebirth/alive"
)

type aliver struct {
	count int
}

func (a *aliver) IsAlive(addr string) (bool, error) {
	pinger, err := goping.NewPinger(addr)
	if err != nil {
		return false, err
	}
	pinger.Count = a.count
	err = pinger.Run()
	if err != nil {
		return false, err
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		return true, nil
	}
	return false, nil
}

func New(count int) (alive.Aliver, error) {
	return &aliver{
		count: count,
	}, nil
}
