package gossip

import "time"

type Config struct {
	Bind           Addr
	Seeds          []Addr
	PingInvterval  time.Duration
	PingTimeout    time.Duration
	SuspectTimeout time.Duration
	MaxPiggyBack   int
	IndirectK      int
}

func DefaultConfig(bind Addr) Config {
	return Config{
		Bind:           bind,
		PingInvterval:  1 * time.Second,
		PingTimeout:    200 * time.Millisecond,
		SuspectTimeout: 5 * time.Second,
		MaxPiggyBack:   16,
		IndirectK:      3,
	}
}
