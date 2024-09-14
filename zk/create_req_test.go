package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCreateReq(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "0000000100000001000000082f7a6b2d746573740000000568656c6c6f000000010000001f00000005776f726c6400000006616e796f6e6500000000")
	req, err := DecodeCreateReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_CREATE, req.OpCode)
	assert.Equal(t, "/zk-test", req.Path)
	assert.Equal(t, []byte("hello"), req.Data)
	assert.Len(t, req.Permissions, 1)
	assert.Equal(t, 31, req.Permissions[0])
	assert.Equal(t, "world", req.Scheme)
	assert.Equal(t, "anyone", req.Credentials)
	assert.Equal(t, 0, req.Flags)
}

func TestEncodeCreateReq(t *testing.T) {
	req := &CreateReq{
		TransactionId: 1,
		OpCode:        OP_CREATE,
		Path:          "/zk-test",
		Data:          []byte("hello"),
		Permissions:   []int{31},
		Scheme:        "world",
		Credentials:   "anyone",
		Flags:         0,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testx.Hex2Bytes(t, "0000000100000001000000082f7a6b2d746573740000000568656c6c6f000000010000001f00000005776f726c6400000006616e796f6e6500000000"), bytes)
}
