package zk

import (
	"github.com/shoothzj/gox/netx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClientGetChildrenData(t *testing.T) {
	config := &Config{
		Address: netx.Address{
			Host: "localhost",
			Port: 2181,
		},
	}
	client, err := NewClient(config)
	require.NoError(t, err)
	defer client.Close()
	getChildrenResp, err := client.GetChildren("/")
	require.Nil(t, err)
	assert.NotNil(t, getChildrenResp)
	getDataResp, err := client.GetData("/zookeeper")
	require.Nil(t, err)
	assert.NotNil(t, getDataResp)
}
