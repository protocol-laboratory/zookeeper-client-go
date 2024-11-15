package zk

import "runtime/debug"

type GetChildrenReq struct {
	TransactionId int32
	OpCode        OpCode
	Path          string
	Watch         bool
}

func DecodeGetChildrenReq(bytes []byte) (req *GetChildrenReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &GetChildrenReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Watch, idx = readWatch(bytes, idx)
	return req, nil
}

func (g *GetChildrenReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(g.Path) + LenWatch
	return length
}

func (g *GetChildrenReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, g.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, g.TransactionId)
	idx = putOpCode(bytes, idx, OpGetChildren)
	idx = putPath(bytes, idx, g.Path)
	idx = putWatch(bytes, idx, g.Watch)
	return bytes
}
