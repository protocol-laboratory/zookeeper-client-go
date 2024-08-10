package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeDeleteResp(t *testing.T) {
	bytes := testHex2Bytes(t, "00000002000000000000002600000000")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 2, resp.TransactionId)
	assert.Equal(t, int64(38), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
}

func TestEncodeDeleteResp(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 2,
		ZxId:          38,
		Error:         EC_OK,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "00000002000000000000002600000000"), bytes)
}

func TestDecodeDeleteRespNoNodeExist(t *testing.T) {
	bytes := testHex2Bytes(t, "000000010000000000000025ffffff9b")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(37), resp.ZxId)
	assert.Equal(t, EC_NoNodeError, resp.Error)
}

func TestEncodeDeleteRespNoNodeExist(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 1,
		ZxId:          37,
		Error:         EC_NoNodeError,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "000000010000000000000025ffffff9b"), bytes)
}
