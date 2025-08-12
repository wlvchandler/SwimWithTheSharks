package gossip

import "time"

type Addr = string

type Config struct {
	Bind           Addr
	Seeds          []Addr
	PingInvterval  time.Duration
	PingTimeout    time.Duration
	SuspectTimeout time.Duration
	MaxPiggyBack   int
	IndirectK      int
}
