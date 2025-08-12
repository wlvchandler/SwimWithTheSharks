package gossip

import "testing"

func TestDefaults(t *testing.T) {
	cfg := DefaultConfig(":7946")
	if cfg.Bind != ":7946" {
		t.Fatalf("expected Bind to be ':7946', got '%q'", cfg.Bind)
	}
}
