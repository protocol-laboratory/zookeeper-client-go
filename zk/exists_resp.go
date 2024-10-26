package zk

import "runtime/debug"

type ExistsResp struct {
	TransactionId int32
	ZxId          int64
	Error         ErrorCode
	Stat          *Stat
}

func DecodeExistsResp(bytes []byte) (resp *ExistsResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &ExistsResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EC_OK {
		resp.Stat, idx = readStat(bytes, idx)
	}
	return resp, nil
}

func (c *ExistsResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EC_OK {
		length += c.Stat.BytesLength()
	}
	return length
}

func (c *ExistsResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EC_OK {
		idx = putStat(bytes, idx, c.Stat)
	}
	return bytes
}
