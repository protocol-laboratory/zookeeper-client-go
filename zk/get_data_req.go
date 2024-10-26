package zk

import "runtime/debug"

type GetDataReq struct {
	TransactionId int32
	OpCode        OpCode
	Path          string
	Watch         bool
}

func DecodeGetDataReq(bytes []byte) (req *GetDataReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &GetDataReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Watch, idx = readWatch(bytes, idx)
	return req, nil
}

func (e *GetDataReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(e.Path) + LenWatch
	return length
}

func (e *GetDataReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, e.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, e.TransactionId)
	idx = putOpCode(bytes, idx, OP_GET_DATA)
	idx = putPath(bytes, idx, e.Path)
	idx = putWatch(bytes, idx, e.Watch)
	return bytes
}
