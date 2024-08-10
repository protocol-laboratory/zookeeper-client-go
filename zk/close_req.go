package zk

import "runtime/debug"

type CloseReq struct {
	TransactionId int
	OpCode        OpCode
}

func DecodeCloseReq(bytes []byte) (req *CloseReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &CloseReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	return req, nil
}

func (c *CloseReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode
	return length
}

func (c *CloseReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putOpCode(bytes, idx, OP_CLOSE_SESSION)
	return bytes
}
