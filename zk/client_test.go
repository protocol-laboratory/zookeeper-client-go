package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCliConnect(t *testing.T) {
	ZkClientConfig := Config{
		Host: "localhost",
		Port: 2181,
	}
	zknetClient, err := NewClient(ZkClientConfig)
	if err != nil {
		t.Error(err)
	}
	defer zknetClient.Close()
	req := &ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         30_000,
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	}
	resp, err := zknetClient.Connect(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, resp)
}
