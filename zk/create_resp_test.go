package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCreateResp(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "00000001000000000000000500000000000000082f7a6b2d74657374")
	resp, err := DecodeCreateResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(5), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
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
	assert.Equal(t, testx.Hex2Bytes(t, "00000001000000000000000500000000000000082f7a6b2d74657374"), bytes)
}

func TestDecodeCreateRespNodeExistsError(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "000000020000000000000020ffffff92")
	resp, err := DecodeCreateResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 2, resp.TransactionId)
	assert.Equal(t, int64(32), resp.ZxId)
	assert.Equal(t, EC_NodeExistsError, resp.Error)
}

func TestEncodeCreateRespNodeExistsError(t *testing.T) {
	resp := &CreateResp{
		TransactionId: 2,
		ZxId:          32,
		Error:         EC_NodeExistsError,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testx.Hex2Bytes(t, "000000020000000000000020ffffff92"), bytes)
}
