package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCloseReq(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "00000003fffffff5")
	req, err := DecodeCloseReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 3, req.TransactionId)
	assert.Equal(t, OP_CLOSE_SESSION, req.OpCode)
}

func TestEncodeCloseReq(t *testing.T) {
	req := &CloseReq{
		TransactionId: 3,
		OpCode:        OP_CLOSE_SESSION,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testx.Hex2Bytes(t, "00000003fffffff5"), bytes)
}
