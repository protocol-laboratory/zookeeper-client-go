package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeGetDataResp(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "0000000300000000000000320000000000000005776f726c640000000000000031000000000000003200000182ef427d6100000182ef427d65000000010000000000000000000000000000000000000005000000000000000000000031")
	resp, err := DecodeGetDataResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 3, resp.TransactionId)
	assert.Equal(t, int64(50), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
	assert.Equal(t, int64(49), resp.Stat.CreatedZxId)
	assert.Equal(t, int64(50), resp.Stat.LastModifiedZxId)
	assert.Equal(t, int64(1661871488353), resp.Stat.Created)
	assert.Equal(t, int64(1661871488357), resp.Stat.LastModified)
	assert.Equal(t, 1, resp.Stat.Version)
	assert.Equal(t, 0, resp.Stat.ChildVersion)
	assert.Equal(t, 0, resp.Stat.AclVersion)
	assert.Equal(t, int64(0), resp.Stat.EphemeralOwner)
	assert.Equal(t, 5, resp.Stat.DataLength)
	assert.Equal(t, 0, resp.Stat.NumChildren)
	assert.Equal(t, int64(49), resp.Stat.LastModifiedChildrenZxId)
}

func TestEncodeGetDataResp(t *testing.T) {
	resp := &GetDataResp{
		TransactionId: 3,
		ZxId:          50,
		Error:         EC_OK,
		Data:          []byte("world"),
		Stat: &Stat{
			CreatedZxId:              49,
			LastModifiedZxId:         50,
			Created:                  1661871488353,
			LastModified:             1661871488357,
			Version:                  1,
			ChildVersion:             0,
			AclVersion:               0,
			EphemeralOwner:           0,
			DataLength:               5,
			NumChildren:              0,
			LastModifiedChildrenZxId: 49,
		},
	}
	bytes := resp.Bytes()
	assert.Equal(t, testx.Hex2Bytes(t, "0000000300000000000000320000000000000005776f726c640000000000000031000000000000003200000182ef427d6100000182ef427d65000000010000000000000000000000000000000000000005000000000000000000000031"), bytes)
}
