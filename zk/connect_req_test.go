package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeConnectReq(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "000000000000000000000000000075300000000000000000000000100000000000000000000000000000000000")
	req, err := DecodeConnectReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 0, req.ProtocolVersion)
	assert.Equal(t, int64(0), req.LastZxidSeen)
	assert.Equal(t, 30_000, req.Timeout)
	assert.Equal(t, int64(0), req.SessionId)
	assert.Len(t, req.Password, 16)
}

func TestEncodeConnectReq(t *testing.T) {
	req := &ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         30_000,
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testx.Hex2Bytes(t, "000000000000000000000000000075300000000000000000000000100000000000000000000000000000000000"), bytes)
}
