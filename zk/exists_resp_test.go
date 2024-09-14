package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeExistsResp(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "00000001000000000000001f00000000000000000000001c000000000000001d00000182ee57606900000182ee57606c00000001000000000000000000000000000000000000000500000000000000000000001c")
	resp, err := DecodeExistsResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(31), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
	assert.Equal(t, int64(28), resp.Stat.CreatedZxId)
	assert.Equal(t, int64(29), resp.Stat.LastModifiedZxId)
	assert.Equal(t, int64(1661856079977), resp.Stat.Created)
	assert.Equal(t, int64(1661856079980), resp.Stat.LastModified)
	assert.Equal(t, 1, resp.Stat.Version)
	assert.Equal(t, 0, resp.Stat.ChildVersion)
	assert.Equal(t, 0, resp.Stat.AclVersion)
	assert.Equal(t, int64(0), resp.Stat.EphemeralOwner)
	assert.Equal(t, 5, resp.Stat.DataLength)
	assert.Equal(t, 0, resp.Stat.NumChildren)
	assert.Equal(t, int64(28), resp.Stat.LastModifiedChildrenZxId)
}

func TestEncodeExistsResp(t *testing.T) {
	resp := &ExistsResp{
		TransactionId: 1,
		ZxId:          31,
		Error:         EC_OK,
		Stat: &Stat{
			CreatedZxId:              28,
			LastModifiedZxId:         29,
			Created:                  1661856079977,
			LastModified:             1661856079980,
			Version:                  1,
			ChildVersion:             0,
			AclVersion:               0,
			EphemeralOwner:           0,
			DataLength:               5,
			NumChildren:              0,
			LastModifiedChildrenZxId: 28,
		},
	}
	bytes := resp.Bytes()
	assert.Equal(t, testx.Hex2Bytes(t, "00000001000000000000001f00000000000000000000001c000000000000001d00000182ee57606900000182ee57606c00000001000000000000000000000000000000000000000500000000000000000000001c"), bytes)
}

func TestDecodeExistsRespNoNodeExist(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "00000001000000000000001bffffff9b")
	resp, err := DecodeExistsResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(27), resp.ZxId)
	assert.Equal(t, EC_NoNodeError, resp.Error)
}

func TestEncodeExistsRespNoNodeExist(t *testing.T) {
	resp := &ExistsResp{
		TransactionId: 1,
		ZxId:          27,
		Error:         EC_NoNodeError,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testx.Hex2Bytes(t, "00000001000000000000001bffffff9b"), bytes)
}
