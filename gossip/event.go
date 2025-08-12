package gossip

type EventType uint8 // emitted on Node.Events() TODO

const (
	EventJoin EventType = iota
	EventAlive
	EventUpdate
	EventSuspect
	EventDead
	EventLeave
)

type Event struct {
	Type   EventType
	Member Member
}
