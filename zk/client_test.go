package zk

import (
	"github.com/shoothzj/gox/netx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCliConnect(t *testing.T) {
	config := Config{
		Address: netx.Address{
			Host: "localhost",
			Port: 2181,
		},
	}
	client, err := NewClient(config)
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
