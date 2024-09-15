package zk

import (
	"crypto/tls"
	"fmt"
	"github.com/shoothzj/gox/buffer"
	"github.com/shoothzj/gox/netx"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	Address          netx.Address
	SendQueueSize    int
	PendingQueueSize int
	BufferMax        int
	Timeout          time.Duration
	TlsConfig        *tls.Config
}

type Client struct {
	config       *Config
	client       *ProtocolClient
	conn         net.Conn
	eventsChan   chan *sendRequest
	pendingQueue chan *sendRequest
	buffer       *buffer.Buffer
	closeCh      chan struct{}

	transactionId atomic.Int32
}

func (c *Client) Create(path string, data []byte, permissions []int, scheme string, credentials string, flags int) (*CreateResp, error) {
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
	req := &DeleteReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_DELETE
	req.Path = path
	req.Version = version
	return c.client.Delete(req)
}

func (c *Client) Exists(path string) (*ExistsResp, error) {
	req := &ExistsReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_EXISTS
	req.Path = path
	req.Watch = false
	return c.client.Exists(req)
}

func (c *Client) GetData(path string) (*GetDataResp, error) {
	req := &GetDataReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_GET_DATA
	req.Path = path
	req.Watch = false
	return c.client.GetData(req)
}

func (c *Client) SetData(path string, data []byte, version int) (*SetDataResp, error) {
	req := &SetDataReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_SET_DATA
	req.Path = path
	req.Data = data
	req.Version = version
	return c.client.SetData(req)
}

func (c *Client) GetChildren(path string) (*GetChildrenResp, error) {
	req := &GetChildrenReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_GET_CHILDREN
	req.Path = path
	req.Watch = false
	return c.client.GetChildren(req)
}

func (c *Client) CloseSession() (*CloseResp, error) {
	req := &CloseReq{}
	req.TransactionId = c.nextTransactionId()
	req.OpCode = OP_CLOSE_SESSION
	return c.client.CloseSession(req)
}

func (c *Client) nextTransactionId() int {
	return int(c.transactionId.Add(1))
}

func (c *Client) Send(bytes []byte) ([]byte, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var result []byte
	var err error
	c.sendAsync(bytes, func(resp []byte, e error) {
		result = resp
		err = e
		wg.Done()
	})
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) sendAsync(bytes []byte, callback func([]byte, error)) {
	sr := &sendRequest{
		bytes:    bytes,
		callback: callback,
	}
	c.eventsChan <- sr
}

func (c *Client) read() {
	for {
		select {
		case req := <-c.pendingQueue:
			n, err := c.conn.Read(c.buffer.WritableSlice())
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			err = c.buffer.AdjustWriteCursor(n)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			if c.buffer.Size() < 4 {
				continue
			}
			bytes := make([]byte, 4)
			err = c.buffer.ReadExactly(bytes)
			c.buffer.Compact()
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			length := int(bytes[3]) | int(bytes[2])<<8 | int(bytes[1])<<16 | int(bytes[0])<<24
			if c.buffer.Size() < length {
				continue
			}
			// in case ddos attack
			if length > c.buffer.Capacity() {
				req.callback(nil, fmt.Errorf("response length %d is too large", length))
				c.closeCh <- struct{}{}
				break
			}
			data := make([]byte, length)
			err = c.buffer.ReadExactly(data)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			c.buffer.Compact()
			req.callback(data, nil)
		case <-c.closeCh:
			return
		}
	}
}

func (c *Client) write() {
	for {
		select {
		case req := <-c.eventsChan:
			n, err := c.conn.Write(req.bytes)
			if err != nil {
				req.callback(nil, err)
				c.closeCh <- struct{}{}
				break
			}
			if n != len(req.bytes) {
				req.callback(nil, fmt.Errorf("write %d bytes, but expect %d bytes", n, len(req.bytes)))
				c.closeCh <- struct{}{}
				break
			}
			c.pendingQueue <- req
		case <-c.closeCh:
			return
		}
	}
}

func (c *Client) Close() {
	c.client.Close()
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

	protocolClient, err := NewProtocolClient(config.Address, config)
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

	client := &Client{
		config: config,
		client: protocolClient,
	}
	return client, nil
}
