package codec

import "github.com/wlvchandler/SwimWithTheSharks/gossip"

const Version uint8 = 1

type Digest struct {
	ID          gossip.NodeID      `json:"id"`
	Incarnation gossip.Incarnation `json:"inc"`
	Status      uint8              `json:"st"`
	Heartbeat   uint64             `json:"hb"`
}

// on the wire message container
type Envelope struct {
	Type string        `json:"t"` // ping/ack/pingreq/join/joinack/leave
	From gossip.NodeID `json:"f"`
	V    uint8         `json:"v"`
	Body jsonRaw       `json:"b"` // raw json body
}

// body types (keep fields minimal, piggyback digests everywhere)
type Ping struct {
	Seq     uint64   `json:"s"`
	Digests []Digest `json:"d,omitempty"`
}

type Ack struct {
	Seq     uint64   `json:"s"`
	Digests []Digest `json:"d,omitempty"`
}

type PingReq struct {
	Seq     uint64      `json:"s"`
	Target  gossip.Addr `json:"tg"`
	Digests []Digest    `json:"d,omitempty"`
}

type Join struct {
	ID   gossip.NodeID `json:"id"`
	Addr gossip.Addr   `json:"addr"`
}

type JoinAck struct {
	Members []Digest `json:"m"`
}

type Leave struct {
	Reason string `json:"r,omitempty"`
}
