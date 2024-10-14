package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCloseResp(t *testing.T) {
	bytes := hex2Bytes(t, "00000003000000000000000700000000")
	resp, err := DecodeCloseResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 3, resp.TransactionId)
	assert.Equal(t, int64(7), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
}

func TestEncodeCloseResp(t *testing.T) {
	resp := &CloseResp{
		TransactionId: 3,
		ZxId:          int64(7),
		Error:         0,
	}
	bytes := resp.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "00000003000000000000000700000000"), bytes)
}
