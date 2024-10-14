package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeGetChildrenResp(t *testing.T) {
	bytes := hex2Bytes(t, "00000001000000000000002200000000000000030000000a65786973742d74657374000000077a6b2d74657374000000097a6f6f6b6565706572")
	resp, err := DecodeGetChildrenResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(34), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
	assert.Equal(t, 3, len(resp.Children))
	assert.Equal(t, "exist-test", resp.Children[0])
	assert.Equal(t, "zk-test", resp.Children[1])
	assert.Equal(t, "zookeeper", resp.Children[2])
}

func TestEncodeGetChildrenResp(t *testing.T) {
	resp := &GetChildrenResp{
		TransactionId: 1,
		ZxId:          34,
		Error:         EC_OK,
		Children:      []string{"exist-test", "zk-test", "zookeeper"},
	}
	bytes := resp.Bytes()
	assert.Equal(t, hex2Bytes(t, "00000001000000000000002200000000000000030000000a65786973742d74657374000000077a6b2d74657374000000097a6f6f6b6565706572"), bytes)
}
