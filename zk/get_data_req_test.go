package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeGetDataReq(t *testing.T) {
	bytes := hex2Bytes(t, "0000000300000004000000082f7a6b2d7465737400")
	req, err := DecodeGetDataReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 3, req.TransactionId)
	assert.Equal(t, OP_GET_DATA, req.OpCode)
	assert.Equal(t, "/zk-test", req.Path)
	assert.False(t, req.Watch)
}

func TestEncodeGetDataReq(t *testing.T) {
	req := &GetDataReq{
		TransactionId: 3,
		OpCode:        OP_GET_DATA,
		Path:          "/zk-test",
		Watch:         false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "0000000300000004000000082f7a6b2d7465737400"), bytes)
}
