package p2p

import "errors"

type HandshakeFunc func(any) error

// ErrInvalidHandshake is returned if the handshake between the local and remote node could not be established
var ErrInvalidHandshake = errors.New("invalid handshake")

func NOPHandshakeFunc(any) error { return nil }
