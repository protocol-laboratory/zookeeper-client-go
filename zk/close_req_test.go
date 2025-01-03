package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCloseReq(t *testing.T) {
	bytes := hex2Bytes(t, "00000003fffffff5")
	req, err := DecodeCloseReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(3), req.TransactionId)
	assert.Equal(t, OpCloseSession, req.OpCode)
}

func TestEncodeCloseReq(t *testing.T) {
	req := &CloseReq{
		TransactionId: 3,
		OpCode:        OpCloseSession,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "00000003fffffff5"), bytes)
}
