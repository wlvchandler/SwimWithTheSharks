package gossip

import (
	"testing"
	"time"
)

func TestDefaultsAndValidation(t *testing.T) {
	cfg := DefaultConfig(":7946").WithDefaults()
	if err := cfg.Validate(); err != nil {
		t.Fatalf("validate default: %v", err)
	}
	if cfg.Advertise != ":7946" {
		t.Fatalf("advertise default: %q", cfg.Advertise)
	}
	if cfg.ID == "" {
		t.Fatalf("ID not generated")
	}
	if cfg.PingTimeout >= cfg.PingInterval {
		t.Fatalf("PingTimeout should be < PingInterval")
	}
	if cfg.SuspectTimeout < 3*cfg.PingInterval {
		t.Fatalf("SuspectTimeout should be >= 3*PingInterval")
	}
	if cfg.MaxPiggyBack <= 0 || cfg.IndirectK < 0 {
		t.Fatalf("bad defaults: %v %v", cfg.MaxPiggyBack, cfg.IndirectK)
	}
}

func TestValidateFailures(t *testing.T) {
	bad := DefaultConfig("").WithDefaults()
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected Bind error")
	}
	bad = DefaultConfig(":x").WithDefaults()
	bad.PingInterval = 200 * time.Millisecond
	bad.PingTimeout = 300 * time.Millisecond
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected PingTimeout < PingInterval error")
	}
	bad = DefaultConfig(":7946").WithDefaults()
	bad.SuspectTimeout = bad.PingInterval * 2
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected SuspectTimeout >= 3*PingInterval error")
	}
}
