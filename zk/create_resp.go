package zk

import "runtime/debug"

type CreateResp struct {
	TransactionId int
	ZxId          int64
	Error         ErrorCode
	Path          string
}

func DecodeCreateResp(bytes []byte) (resp *CreateResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &CreateResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EC_OK {
		resp.Path, idx = readPath(bytes, idx)
	}
	return resp, nil
}

func (c *CreateResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EC_OK {
		length += StrLen(c.Path)
	}
	return length
}

func (c *CreateResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EC_OK {
		idx = putPath(bytes, idx, c.Path)
	}
	return bytes
}
