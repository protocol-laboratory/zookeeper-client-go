package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeDeleteResp(t *testing.T) {
	bytes := hex2Bytes(t, "00000002000000000000002600000000")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), resp.TransactionId)
	assert.Equal(t, int64(38), resp.ZxId)
	assert.Equal(t, EcOk, resp.Error)
}

func TestEncodeDeleteResp(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 2,
		ZxId:          38,
		Error:         EcOk,
	}
	bytes := resp.Bytes()
	assert.Equal(t, hex2Bytes(t, "00000002000000000000002600000000"), bytes)
}

func TestDecodeDeleteRespNoNodeExist(t *testing.T) {
	bytes := hex2Bytes(t, "000000010000000000000025ffffff9b")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), resp.TransactionId)
	assert.Equal(t, int64(37), resp.ZxId)
	assert.Equal(t, EcNoNode, resp.Error)
}

func TestEncodeDeleteRespNoNodeExist(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 1,
		ZxId:          37,
		Error:         EcNoNode,
	}
	bytes := resp.Bytes()
	assert.Equal(t, hex2Bytes(t, "000000010000000000000025ffffff9b"), bytes)
}
