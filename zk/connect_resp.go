package zk

import "runtime/debug"

type ConnectResp struct {
	ProtocolVersion int
	Timeout         int
	SessionId       int64
	Password        []byte
	ReadOnly        bool
}

func DecodeConnectResp(bytes []byte) (resp *ConnectResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &ConnectResp{}
	idx := 0
	resp.ProtocolVersion, idx = readProtocolVersion(bytes, idx)
	resp.Timeout, idx = readTimeout(bytes, idx)
	resp.SessionId, idx = readSessionId(bytes, idx)
	resp.Password, idx = readPassword(bytes, idx)
	resp.ReadOnly, idx = readReadOnly(bytes, idx)
	return resp, nil
}

func (c *ConnectResp) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenProtocolVersion + LenTimeout + LenSessionId + BytesLen(c.Password) + LenReadonly
	return length
}

func (c *ConnectResp) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putProtocolVersion(bytes, idx, c.ProtocolVersion)
	idx = putTimeout(bytes, idx, c.Timeout)
	idx = putSessionId(bytes, idx, c.SessionId)
	idx = putPassword(bytes, idx, c.Password)
	idx = putBool(bytes, idx, c.ReadOnly)
	return bytes
}
