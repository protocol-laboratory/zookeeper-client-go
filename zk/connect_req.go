package zk

import "runtime/debug"

type ConnectReq struct {
	ProtocolVersion int
	LastZxidSeen    int64
	Timeout         int
	SessionId       int64
	Password        []byte
	ReadOnly        bool
}

func DecodeConnectReq(bytes []byte) (req *ConnectReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &ConnectReq{}
	idx := 0
	req.ProtocolVersion, idx = readProtocolVersion(bytes, idx)
	req.LastZxidSeen, idx = readLastZxidSeen(bytes, idx)
	req.Timeout, idx = readTimeout(bytes, idx)
	req.SessionId, idx = readSessionId(bytes, idx)
	req.Password, idx = readPassword(bytes, idx)
	req.ReadOnly, idx = readReadOnly(bytes, idx)
	return req, nil
}

func (c *ConnectReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenProtocolVersion + LenLastZxidSeen + LenTimeout + LenSessionId + BytesLen(c.Password) + LenReadonly
	return length
}

func (c *ConnectReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putProtocolVersion(bytes, idx, c.ProtocolVersion)
	idx = putLastZxidSeen(bytes, idx, c.LastZxidSeen)
	idx = putTimeout(bytes, idx, c.Timeout)
	idx = putSessionId(bytes, idx, c.SessionId)
	idx = putPassword(bytes, idx, c.Password)
	idx = putBool(bytes, idx, c.ReadOnly)
	return bytes
}
