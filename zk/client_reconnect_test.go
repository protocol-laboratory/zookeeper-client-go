package zk

import (
	"github.com/shoothzj/gox/netx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientReconnect(t *testing.T) {
	config := &Config{
		Addresses: []netx.Address{
			{
				Host: "localhost",
				Port: 2181,
			},
		},
	}
	client, err := NewClient(config)
	require.NoError(t, err)
	getChildrenResp, err := client.GetChildren("/")
	require.Nil(t, err)
	assert.NotNil(t, getChildrenResp)
	tempProtocolClient := client.client
	tempProtocolClient.Close()
	getChildrenResp, err = client.GetChildren("/")
	require.Equal(t, ErrClientClosed, err)
	time.Sleep(time.Second * 3)
	client.mutex.RLock()
	assert.NotEqual(t, tempProtocolClient, client.client)
	client.mutex.RUnlock()
	getChildrenResp, err = client.GetChildren("/")
	require.Nil(t, err)
	assert.NotNil(t, getChildrenResp)
}
