package zk

import "errors"

var (
	ErrClientClosed = errors.New("client state is closed")
)
