package gossip

import (
	"log/slog"
	"math/rand"
	"time"
)

type Config struct {
	// networking
	Bind      Addr // local udp bind. e.g., ":7946" or "0.0.0.0:7946"
	Advertise Addr // defaults to Bind. what peers should dial
	Seeds     []Addr

	// timers
	PingInterval   time.Duration
	PingTimeout    time.Duration
	SuspectTimeout time.Duration

	// swim parameters
	MaxPiggyBack int // cap on digests piggybacked per message
	IndirectK    int // number of helpers for ping-req

	// infra
	Log   *slog.Logger
	Rand  *rand.Rand
	Clock Clock

	// identity (optional. if empty we auto-gen a random one)
	ID NodeID
}

func DefaultConfig(bind Addr) Config {
	return Config{
		Bind:           bind,
		Advertise:      "", // resolved at start
		PingInterval:   1 * time.Second,
		PingTimeout:    200 * time.Millisecond,
		SuspectTimeout: 5 * time.Second,
		MaxPiggyBack:   16,
		IndirectK:      3,
	}
}

func (c Config) WithDefaults() Config {
	out := c

	if out.Advertise == "" {
		out.Advertise = out.Bind
	}

	if out.PingInterval <= 0 {
		out.PingInterval = 1 * time.Second
	}
}
