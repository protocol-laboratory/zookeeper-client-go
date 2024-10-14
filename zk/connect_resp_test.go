package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeConnectResp(t *testing.T) {
	bytes := hex2Bytes(t, "00000000000075300100020abac20001000000109f9f123fc926a4013f35eca7ee76386100")
	resp, err := DecodeConnectResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 0, resp.ProtocolVersion)
	assert.Equal(t, 30_000, resp.Timeout)
	assert.Equal(t, int64(72059839144132609), resp.SessionId)
	assert.Len(t, resp.Password, 16)
}

func TestEncodeConnectResp(t *testing.T) {
	resp := &ConnectResp{
		ProtocolVersion: 0,
		Timeout:         30_000,
		SessionId:       72059839144132609,
		Password:        hex2Bytes(t, "9f9f123fc926a4013f35eca7ee763861"),
		ReadOnly:        false,
	}
	bytes := resp.Bytes(false)
	assert.Equal(t, hex2Bytes(t, "00000000000075300100020abac20001000000109f9f123fc926a4013f35eca7ee76386100"), bytes)
}
