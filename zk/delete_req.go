package zk

import "runtime/debug"

type DeleteReq struct {
	TransactionId int32
	OpCode        OpCode
	Path          string
	Version       int
}

func DecodeDeleteReq(bytes []byte) (req *DeleteReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &DeleteReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Version, idx = readVersion(bytes, idx)
	return req, nil
}

func (e *DeleteReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(e.Path) + LenVersion
	return length
}

func (e *DeleteReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, e.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, e.TransactionId)
	idx = putOpCode(bytes, idx, OpDelete)
	idx = putPath(bytes, idx, e.Path)
	idx = putVersion(bytes, idx, e.Version)
	return bytes
}
