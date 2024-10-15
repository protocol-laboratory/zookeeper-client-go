package zk

import (
	"crypto/tls"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libgox/addr"
)

type Config struct {
	Addresses        []addr.Address
	SendQueueSize    int
	PendingQueueSize int
	BufferMax        int
	Timeout          time.Duration
	TlsConfig        *tls.Config
}

type Client struct {
	config *Config

	client *ProtocolClient
	mutex  sync.RWMutex

	transactionId atomic.Int32

	reconnectCh              chan time.Time
	lastClientConnectSuccess atomic.Value
}

func (c *Client) Create(path string, data []byte, permissions []int, scheme string, credentials string, flags int) (*CreateResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &CreateReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_CREATE
	req.Path = path
	req.Data = data
	req.Permissions = permissions
	req.Scheme = scheme
	req.Credentials = credentials
	req.Flags = flags
	return c.client.Create(req)
}

func (c *Client) Delete(path string, version int) (*DeleteResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &DeleteReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_DELETE
	req.Path = path
	req.Version = version
	return c.client.Delete(req)
}

func (c *Client) Exists(path string) (*ExistsResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &ExistsReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_EXISTS
	req.Path = path
	req.Watch = false
	return c.client.Exists(req)
}

func (c *Client) GetData(path string) (*GetDataResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &GetDataReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_GET_DATA
	req.Path = path
	req.Watch = false
	return c.client.GetData(req)
}

func (c *Client) SetData(path string, data []byte, version int) (*SetDataResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &SetDataReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_SET_DATA
	req.Path = path
	req.Data = data
	req.Version = version
	return c.client.SetData(req)
}

func (c *Client) GetChildren(path string) (*GetChildrenResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &GetChildrenReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_GET_CHILDREN
	req.Path = path
	req.Watch = false
	return c.client.GetChildren(req)
}

func (c *Client) CloseSession() (*CloseResp, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	req := &CloseReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_CLOSE_SESSION
	return c.client.CloseSession(req)
}

func (c *Client) nextTransactionId() int {
	return int(c.transactionId.Add(1))
}

func (c *Client) Close() {
	close(c.reconnectCh)
	c.client.Close()
}

func (c *Client) reconnect() {
	for timestamp := range c.reconnectCh {
		lastConnect, ok := c.lastClientConnectSuccess.Load().(time.Time)
		if ok {
			if timestamp.Sub(lastConnect) < 0 {
				continue
			}
		}

		c.mutex.Lock()
		// Close the old client if needed
		if c.client != nil {
			c.client.Close()
		}

		// Create a new client
		selectedAddress := c.config.Addresses[rand.Intn(len(c.config.Addresses))]
		newClient, err := NewProtocolClient(selectedAddress, c.config, c.reconnectCh)
		if err != nil {
			c.mutex.Unlock()
			continue
		}

		_, err = newClient.Connect(&ConnectReq{
			ProtocolVersion: 0,
			LastZxidSeen:    0,
			Timeout:         int(c.config.Timeout.Milliseconds()),
			SessionId:       0,
			Password:        PasswordEmpty,
			ReadOnly:        false,
		})
		if err != nil {
			c.mutex.Unlock()
			newClient.Close()
			continue
		}

		// Replace with the new client
		c.client = newClient
		c.mutex.Unlock()
	}
}

func NewClient(config *Config) (*Client, error) {
	if config.SendQueueSize == 0 {
		config.SendQueueSize = 1000
	}
	if config.PendingQueueSize == 0 {
		config.PendingQueueSize = 1000
	}
	if config.BufferMax == 0 {
		config.BufferMax = 512 * 1024
	}

	client := &Client{
		config:      config,
		reconnectCh: make(chan time.Time),
	}

	selectedAddress := config.Addresses[rand.Intn(len(config.Addresses))]

	protocolClient, err := NewProtocolClient(selectedAddress, config, client.reconnectCh)
	if err != nil {
		return nil, err
	}

	_, err = protocolClient.Connect(&ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         int(config.Timeout.Milliseconds()),
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	})
	if err != nil {
		protocolClient.Close()
		return nil, err
	}

	client.client = protocolClient
	client.lastClientConnectSuccess.Store(time.Now())

	go client.reconnect()

	return client, nil
}
