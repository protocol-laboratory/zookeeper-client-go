package zk

import (
	"testing"

	"github.com/libgox/addr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGetChildrenData(t *testing.T) {
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
	defer client.Close()
	getChildrenResp, err := client.GetChildren("/")
	require.Nil(t, err)
	assert.NotNil(t, getChildrenResp)
	getDataResp, err := client.GetData("/zookeeper")
	require.Nil(t, err)
	assert.NotNil(t, getDataResp)
}

func TestClientDoubleClose(t *testing.T) {
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
	client.Close()
	client.Close()
}

func TestClientSendReqAfterProtocolClientCloseNoDeadLock(t *testing.T) {
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
	client.client.Close()
	_, err = client.GetChildren("/")
	_, err = client.GetData("/zookeeper")
}
