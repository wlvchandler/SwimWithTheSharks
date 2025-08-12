package codec

import (
	"encoding/json"
	"errors"
)

type jsonRaw = json.RawMessage

var (
	ErrUnknownType = errors.New("codec: unknown message type")
	ErrBadVersion  = errors.New("codec: unsupported version")
	ErrDecodeBody  = errors.New("codec: body decode failed")
)

//TODO: marshal/unmarshal
