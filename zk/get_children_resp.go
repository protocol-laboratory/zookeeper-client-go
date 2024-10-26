package zk

import "runtime/debug"

type GetChildrenResp struct {
	TransactionId int32
	ZxId          int64
	Error         ErrorCode
	Children      []string
}

func DecodeGetChildrenResp(bytes []byte) (resp *GetChildrenResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &GetChildrenResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EC_OK {
		var length int
		length, idx = readInt(bytes, idx)
		for i := 0; i < length; i++ {
			var child string
			child, idx = readString(bytes, idx)
			resp.Children = append(resp.Children, child)
		}
	}
	return resp, nil
}

func (c *GetChildrenResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EC_OK {
		length += LenArray
		for _, child := range c.Children {
			length += StrLen(child)
		}
	}
	return length
}

func (c *GetChildrenResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EC_OK {
		idx = putInt(bytes, idx, len(c.Children))
		for _, child := range c.Children {
			idx = putString(bytes, idx, child)
		}
	}
	return bytes
}
