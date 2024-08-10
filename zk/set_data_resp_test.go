package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeSetDataResp(t *testing.T) {
	bytes := testHex2Bytes(t, "000000020000000000000006000000000000000000000005000000000000000600000182ec1b539f00000182ec1b53ac000000010000000000000000000000000000000000000005000000000000000000000005")
	resp, err := DecodeSetDataResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 2, resp.TransactionId)
	assert.Equal(t, int64(6), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
	assert.Equal(t, int64(5), resp.Stat.CreatedZxId)
	assert.Equal(t, int64(6), resp.Stat.LastModifiedZxId)
	assert.Equal(t, int64(1661818590111), resp.Stat.Created)
	assert.Equal(t, int64(1661818590124), resp.Stat.LastModified)
	assert.Equal(t, 1, resp.Stat.Version)
	assert.Equal(t, 0, resp.Stat.ChildVersion)
	assert.Equal(t, 0, resp.Stat.AclVersion)
	assert.Equal(t, int64(0), resp.Stat.EphemeralOwner)
	assert.Equal(t, 5, resp.Stat.DataLength)
	assert.Equal(t, 0, resp.Stat.NumChildren)
	assert.Equal(t, int64(5), resp.Stat.LastModifiedChildrenZxId)
}

func TestEncodeSetDataResp(t *testing.T) {
	resp := &SetDataResp{
		TransactionId: 2,
		ZxId:          6,
		Error:         0,
		Stat: &Stat{
			CreatedZxId:              5,
			LastModifiedZxId:         6,
			Created:                  1661818590111,
			LastModified:             1661818590124,
			Version:                  1,
			ChildVersion:             0,
			AclVersion:               0,
			EphemeralOwner:           0,
			DataLength:               5,
			NumChildren:              0,
			LastModifiedChildrenZxId: 5,
		},
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "000000020000000000000006000000000000000000000005000000000000000600000182ec1b539f00000182ec1b53ac000000010000000000000000000000000000000000000005000000000000000000000005"), bytes)
}
