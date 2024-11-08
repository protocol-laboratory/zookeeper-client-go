package zk

import "runtime/debug"

type GetDataResp struct {
	TransactionId int32
	ZxId          int64
	Data          []byte
	Error         ErrorCode
	Stat          *Stat
}

func DecodeGetDataResp(bytes []byte) (resp *GetDataResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &GetDataResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EcOk {
		resp.Data, idx = readData(bytes, idx)
		resp.Stat, idx = readStat(bytes, idx)
	}
	return resp, nil
}

func (c *GetDataResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EcOk {
		length += BytesLen(c.Data)
		length += c.Stat.BytesLength()
	}
	return length
}

func (c *GetDataResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EcOk {
		idx = putData(bytes, idx, c.Data)
		idx = putStat(bytes, idx, c.Stat)
	}
	return bytes
}
