package zk

import "runtime/debug"

type SetDataReq struct {
	TransactionId int32
	OpCode        OpCode
	Path          string
	Data          []byte
	Version       int
}

func DecodeSetDataReq(bytes []byte) (req *SetDataReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &SetDataReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Data, idx = readData(bytes, idx)
	req.Version, idx = readVersion(bytes, idx)
	return req, nil
}

func (s *SetDataReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(s.Path) + BytesLen(s.Data) + LenVersion
	return length
}

func (s *SetDataReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, s.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, s.TransactionId)
	idx = putOpCode(bytes, idx, OpSetData)
	idx = putPath(bytes, idx, s.Path)
	idx = putData(bytes, idx, s.Data)
	idx = putVersion(bytes, idx, s.Version)
	return bytes
}
