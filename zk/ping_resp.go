package zk

import "runtime/debug"

type PingResp struct {
	TransactionId int32
}

func DecodePingResp(bytes []byte) (req *PingResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &PingResp{}
	idx := 0
	aux, idx := readTransactionId(bytes, idx)
	req.TransactionId = int32(aux)
	return req, nil
}

func (p *PingResp) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId
	return length
}

func (p *PingResp) Bytes(containLen bool) []byte {
	bytes := make([]byte, p.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, p.TransactionId)
	return bytes
}
