package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeGetChildrenReq(t *testing.T) {
	bytes := hex2Bytes(t, "0000000100000008000000012f00")
	req, err := DecodeGetChildrenReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_GET_CHILDREN, req.OpCode)
	assert.Equal(t, "/", req.Path)
	assert.False(t, req.Watch)
}

func TestEncodeGetChildrenReq(t *testing.T) {
	req := &GetChildrenReq{
		TransactionId: 1,
		OpCode:        OP_GET_CHILDREN,
		Path:          "/",
		Watch:         false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "0000000100000008000000012f00"), bytes)
}
