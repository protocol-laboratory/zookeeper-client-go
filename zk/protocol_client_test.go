package zk

import (
	"testing"
	"time"

	"github.com/libgox/addr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtocolClientConnect(t *testing.T) {
	client, err := NewProtocolClient(addr.Address{
		Host: "localhost",
		Port: 2181,
	}, &Config{
		SendQueueSize:    100,
		PendingQueueSize: 100,
		BufferMax:        1024,
	}, make(chan time.Time))
	require.NoError(t, err)
	defer client.Close()
	req := &ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         30_000,
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	}
	resp, err := client.Connect(req)
	require.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestProtocolClientConnectAfterClose(t *testing.T) {
	reconnectChannel := make(chan time.Time, 1024)
	client, err := NewProtocolClient(addr.Address{
		Host: "localhost",
		Port: 2181,
	}, &Config{
		SendQueueSize:    100,
		PendingQueueSize: 100,
		BufferMax:        1024,
	}, reconnectChannel)
	require.NoError(t, err)
	req := &ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         30_000,
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	}
	resp, err := client.Connect(req)
	require.Nil(t, err)
	assert.NotNil(t, resp)
	client.Close()
	_, err = client.Connect(req)
	assert.Equal(t, ErrClientClosed, err)
	close(reconnectChannel)
}
