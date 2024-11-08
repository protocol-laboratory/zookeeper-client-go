package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCreateResp(t *testing.T) {
	bytes := hex2Bytes(t, "00000001000000000000000500000000000000082f7a6b2d74657374")
	resp, err := DecodeCreateResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), resp.TransactionId)
	assert.Equal(t, int64(5), resp.ZxId)
	assert.Equal(t, EcOk, resp.Error)
	assert.Equal(t, "/zk-test", resp.Path)
}

func TestEncodeCreateResp(t *testing.T) {
	resp := &CreateResp{
		TransactionId: 1,
		ZxId:          5,
		Error:         0,
		Path:          "/zk-test",
	}
	bytes := resp.Bytes()
	assert.Equal(t, hex2Bytes(t, "00000001000000000000000500000000000000082f7a6b2d74657374"), bytes)
}

func TestDecodeCreateRespNodeExistsError(t *testing.T) {
	bytes := hex2Bytes(t, "000000020000000000000020ffffff92")
	resp, err := DecodeCreateResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), resp.TransactionId)
	assert.Equal(t, int64(32), resp.ZxId)
	assert.Equal(t, EcNodeExists, resp.Error)
}

func TestEncodeCreateRespNodeExistsError(t *testing.T) {
	resp := &CreateResp{
		TransactionId: 2,
		ZxId:          32,
		Error:         EcNodeExists,
	}
	bytes := resp.Bytes()
	assert.Equal(t, hex2Bytes(t, "000000020000000000000020ffffff92"), bytes)
}
