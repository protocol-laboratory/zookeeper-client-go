package zk

import (
	"fmt"
	"github.com/shoothzj/gox/buffer"
	"github.com/shoothzj/gox/netx"
	"net"
	"sync"
)

type sendRequest struct {
	bytes    []byte
	callback func([]byte, error)
}

type ProtocolClient struct {
	config       *Config
	conn         net.Conn
	eventsChan   chan *sendRequest
	pendingQueue chan *sendRequest
	buffer       *buffer.Buffer
	closeCh      chan struct{}
}

func (c *ProtocolClient) Connect(req *ConnectReq) (*ConnectResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeConnectResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) Create(req *CreateReq) (*CreateResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeCreateResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) Delete(req *DeleteReq) (*DeleteResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeDeleteResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) Exists(req *ExistsReq) (*ExistsResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeExistsResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) GetData(req *GetDataReq) (*GetDataResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeGetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) SetData(req *SetDataReq) (*SetDataResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeSetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) GetChildren(req *GetChildrenReq) (*GetChildrenResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeGetChildrenResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) CloseSession(req *CloseReq) (*CloseResp, error) {
	bytes, err := c.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeCloseResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProtocolClient) Send(bytes []byte) ([]byte, error) {
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

func (c *ProtocolClient) sendAsync(bytes []byte, callback func([]byte, error)) {
	sr := &sendRequest{
		bytes:    bytes,
		callback: callback,
	}
	c.eventsChan <- sr
}

func (c *ProtocolClient) read() {
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

func (c *ProtocolClient) write() {
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

func (c *ProtocolClient) Close() {
	_ = c.conn.Close()
	c.closeCh <- struct{}{}
}

func NewProtocolClient(address netx.Address, config *Config) (*ProtocolClient, error) {
	conn, err := netx.Dial(address, config.TlsConfig)

	if err != nil {
		return nil, err
	}
	client := &ProtocolClient{
		config:       config,
		conn:         conn,
		eventsChan:   make(chan *sendRequest, config.SendQueueSize),
		pendingQueue: make(chan *sendRequest, config.PendingQueueSize),
		buffer:       buffer.NewBuffer(config.BufferMax),
		closeCh:      make(chan struct{}),
	}
	go func() {
		client.read()
	}()
	go func() {
		client.write()
	}()
	return client, nil
}