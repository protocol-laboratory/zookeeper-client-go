package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeSetDataReq(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "0000000200000005000000082f7a6b2d7465737400000005776f726c64ffffffff")
	req, err := DecodeSetDataReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 2, req.TransactionId)
	assert.Equal(t, OP_SET_DATA, req.OpCode)
	assert.Equal(t, "/zk-test", req.Path)
	assert.Equal(t, []byte("world"), req.Data)
	assert.Equal(t, 4294967295, req.Version)
}

func TestEncodeSetDataReq(t *testing.T) {
	req := &SetDataReq{
		TransactionId: 2,
		OpCode:        OP_SET_DATA,
		Path:          "/zk-test",
		Data:          []byte("world"),
		Version:       4294967295,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testx.Hex2Bytes(t, "0000000200000005000000082f7a6b2d7465737400000005776f726c64ffffffff"), bytes)
}
