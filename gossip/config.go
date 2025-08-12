package gossip

import (
	crand "crypto/rand"
	"encoding/hex"
	"errors"
	"log/slog"
	mrand "math/rand"
	"os"
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
	Rand  *mrand.Rand
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

	if out.PingTimeout <= 0 {
		out.PingTimeout = 200 * time.Millisecond
	}

	if out.SuspectTimeout <= 0 {
		out.SuspectTimeout = 5 * time.Second
	}

	if out.MaxPiggyBack <= 0 {
		out.MaxPiggyBack = 16
	}

	if out.IndirectK <= 0 {
		out.IndirectK = 0
	}

	if out.Log == nil {
		out.Log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	if out.Rand == nil {
		var seed int64
		var b [8]byte
		if _, err := crand.Read(b[:]); err == nil {
			seed = int64(uint64(b[0]<<56) | uint64(b[1]<<48) | uint64(b[2]<<40) | uint64(b[3]<<32) |
				uint64(b[4]<<24) | uint64(b[5]<<16) | uint64(b[6]<<8) | uint64(b[7]))
		} else {
			seed = time.Now().UnixNano()
		}
		out.Rand = mrand.New(mrand.NewSource(seed))
	}
	if out.Clock == nil {
		out.Clock = systemClock{}
	}
	if out.ID == "" {
		out.ID = NewNodeID()
	}
	return out
}

func (c Config) Validate() error {
	if c.Bind == "" {
		return errors.New("bind must be set, e.g. ':7946'")
	}
	if c.PingInterval <= 0 {
		return errors.New("PingInterval must be > 0")
	}
	if c.PingTimeout <= 0 {
		return errors.New("PingTimeout must be > 0")
	}
	if c.SuspectTimeout <= 0 {
		return errors.New("SuspectTimeout must be > 0")
	}
	if c.PingTimeout >= c.PingInterval {
		return errors.New("PingTimeout must be < PingInterval")
	}
	if c.SuspectTimeout < 3*c.PingInterval {
		return errors.New("SuspectTimeout should be >= 3 * PingInterval")
	}
	if c.MaxPiggyBack <= 0 {
		return errors.New("MaxPiggyback must be > 0")
	}
	if c.IndirectK < 0 {
		return errors.New("IndirectK must be >= 0")
	}
	return nil
}

func NewNodeID() NodeID {
	var b [16]byte
	_, _ = crand.Read(b[:])
	return NodeID(hex.EncodeToString(b[:]))
}
