package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodePingReq(t *testing.T) {
	bytes := hex2Bytes(t, "fffffffe0000000b")
	req, err := DecodePingReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(-2), req.TransactionId)
	assert.Equal(t, OP_PING, req.OpCode)
}
