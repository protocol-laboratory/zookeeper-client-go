package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeDeleteReq(t *testing.T) {
	bytes := testHex2Bytes(t, "00000001000000020000000c2f7a6b2d6e6f74666f756e6400000000")
	req, err := DecodeDeleteReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_DELETE, req.OpCode)
	assert.Equal(t, "/zk-notfound", req.Path)
	assert.Equal(t, 0, req.Version)
}

func TestEncodeDeleteReq(t *testing.T) {
	req := &DeleteReq{
		TransactionId: 1,
		OpCode:        OP_DELETE,
		Path:          "/zk-notfound",
		Version:       0,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "00000001000000020000000c2f7a6b2d6e6f74666f756e6400000000"), bytes)
}
