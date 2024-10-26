package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodePingResp(t *testing.T) {
	bytes := hex2Bytes(t, "fffffffe000000000000003200000000")
	req, err := DecodePingResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(-2), req.TransactionId)
}
