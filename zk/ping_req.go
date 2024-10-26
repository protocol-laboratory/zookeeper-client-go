package zk

import "runtime/debug"

type PingReq struct {
	TransactionId int32
	OpCode        OpCode
}

func DecodePingReq(bytes []byte) (req *PingReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &PingReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	return req, nil
}

func (p *PingReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode
	return length
}

func (p *PingReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, p.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, p.TransactionId)
	idx = putOpCode(bytes, idx, OP_PING)
	return bytes
}
