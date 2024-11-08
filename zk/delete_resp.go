package zk

import "runtime/debug"

type DeleteResp struct {
	TransactionId int32
	ZxId          int64
	Error         ErrorCode
}

func DecodeDeleteResp(bytes []byte) (resp *DeleteResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &DeleteResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EcOk {
	}
	return resp, nil
}

func (c *DeleteResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EcOk {
	}
	return length
}

func (c *DeleteResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EcOk {
	}
	return bytes
}
