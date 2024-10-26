package zk

import "runtime/debug"

type SetDataResp struct {
	TransactionId int32
	ZxId          int64
	Error         ErrorCode
	Stat          *Stat
}

func DecodeSetDataResp(bytes []byte) (resp *SetDataResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &SetDataResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	resp.Stat, idx = readStat(bytes, idx)
	return resp, nil
}

func (s *SetDataResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError + LenCreatedZxId + LenLastModifiedZxId
	length += LenCreated + LenLastModified + LenVersion + LenChildVersion + LenAclVersion + LenEphemeralOwner
	length += LenDataLength + LenNumberOfChildren + LenLastModifiedZxId
	return length
}

func (s *SetDataResp) Bytes() []byte {
	bytes := make([]byte, s.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, s.TransactionId)
	idx = putZxId(bytes, idx, s.ZxId)
	idx = putError(bytes, idx, s.Error)
	idx = putStat(bytes, idx, s.Stat)
	return bytes
}
