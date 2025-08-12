package codec

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/wlvchandler/SwimWithTheSharks/gossip"
)

type jsonRaw = json.RawMessage

var (
	ErrUnknownType = errors.New("codec: unknown message type")
	ErrBadVersion  = errors.New("codec: unsupported version")
	ErrDecodeBody  = errors.New("codec: body decode failed")
)

func typeOf(b any) (string, error) {
	switch b.(type) {
	case Ping, *Ping:
		return "ping", nil
	case Ack, *Ack:
		return "ack", nil
	case PingReq, *PingReq:
		return "pingreq", nil
	case Join, *Join:
		return "join", nil
	case JoinAck, *JoinAck:
		return "joinack", nil
	case Leave, *Leave:
		return "leave", nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnknownType, reflect.TypeOf(b))
	}
}

func Marshal(from gossip.NodeID, body any) ([]byte, error) {
	typ, err := typeOf(body)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	env := Envelope{Type: typ, From: from, V: Version, Body: b}
	return json.Marshal(&env)
}
