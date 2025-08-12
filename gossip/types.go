package gossip

import "fmt"

type Addr = string        // e.g. "127.0.0.1:666"
type NodeID = string      // id a member in the cluster
type Incarnation = uint64 // increase when a node restarts or force announces newer state
type Status uint8

const (
	StatusAlive Status = iota
	StatusSuspect
	StatusDead
	StatusLeft
)

func (s Status) String() string {
	switch s {
	case StatusAlive:
		return "alive"
	case StatusSuspect:
		return "suspect"
	case StatusDead:
		return "dead"
	case StatusLeft:
		return "left"
	default:
		return fmt.Sprintf("status(%d)", uint8(s))
	}
}

type Member struct {
	ID          NodeID
	Addr        Addr
	Status      Status
	Incarnation Incarnation
	Heartbeat   uint64
	LastUpdate  int64 // unix nano
}
