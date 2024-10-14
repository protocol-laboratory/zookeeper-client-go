package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeExistsReq(t *testing.T) {
	bytes := hex2Bytes(t, "00000001000000030000000b2f65786973742d7465737400")
	req, err := DecodeExistsReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_EXISTS, req.OpCode)
	assert.Equal(t, "/exist-test", req.Path)
	assert.False(t, req.Watch)
}

func TestEncodeExistsReq(t *testing.T) {
	req := &ExistsReq{
		TransactionId: 1,
		OpCode:        OP_EXISTS,
		Path:          "/exist-test",
		Watch:         false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "00000001000000030000000b2f65786973742d7465737400"), bytes)
}
