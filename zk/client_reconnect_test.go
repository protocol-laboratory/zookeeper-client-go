package zk

import (
	"testing"
	"time"

	"github.com/libgox/addr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientReconnect(t *testing.T) {
	config := &Config{
		Addresses: []addr.Address{
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
