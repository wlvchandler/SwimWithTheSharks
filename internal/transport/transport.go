package transport

import (
	"context"
	"net"
)

type Packet struct {
	From *net.UDPAddr
	Data []byte
}

type Transport interface {
	Listen(ctx context.Context, bind string) (<-chan Packet, error)
	Send(ctx context.Context, to string, b []byte) error
	Close() error
}
