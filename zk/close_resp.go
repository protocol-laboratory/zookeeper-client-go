package zk

import "runtime/debug"

type CloseResp struct {
	TransactionId int
	ZxId          int64
	Error         ErrorCode
}

func DecodeCloseResp(bytes []byte) (resp *CloseResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &CloseResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	return resp, nil
}

func (c *CloseResp) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenZxId + LenError
	return length
}

func (c *CloseResp) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	return bytes
}
