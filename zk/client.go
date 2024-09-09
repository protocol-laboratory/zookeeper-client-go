package zk

import (
	"crypto/tls"
	"fmt"
	"github.com/shoothzj/gox/buffer"
	"net"
	"sync"
)

type ZookeeperClientConfig struct {
	Host             string
	Port             int
	BufferMax        int
	SendQueueSize    int
	PendingQueueSize int
	TLSConfig        *tls.Config
}

func (z ZookeeperClientConfig) addr() string {
	return fmt.Sprintf("%s:%d", z.Host, z.Port)
}

type sendRequest struct {
	bytes    []byte
	callback func([]byte, error)
}

type ZookeeperClient struct {
	conn         net.Conn
	eventsChan   chan *sendRequest
	pendingQueue chan *sendRequest
	buffer       *buffer.Buffer
	closeCh      chan struct{}
}

func (z *ZookeeperClient) Connect(req *ConnectReq) (*ConnectResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeConnectResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) Create(req *CreateReq) (*CreateResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeCreateResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) Delete(req *DeleteReq) (*DeleteResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeDeleteResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) Exists(req *ExistsReq) (*ExistsResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeExistsResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) GetData(req *GetDataReq) (*GetDataResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeGetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) SetData(req *SetDataReq) (*SetDataResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeSetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) GetChildren(req *GetChildrenReq) (*GetChildrenResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeGetChildrenResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) CloseSession(req *CloseReq) (*CloseResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := DecodeCloseResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperClient) Send(bytes []byte) ([]byte, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var result []byte
	var err error
	z.sendAsync(bytes, func(resp []byte, e error) {
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

func (z *ZookeeperClient) sendAsync(bytes []byte, callback func([]byte, error)) {
	sr := &sendRequest{
		bytes:    bytes,
		callback: callback,
	}
	z.eventsChan <- sr
}

func (z *ZookeeperClient) read() {
	for {
		select {
		case req := <-z.pendingQueue:
			n, err := z.conn.Read(z.buffer.WritableSlice())
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			err = z.buffer.AdjustWriteCursor(n)
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			if z.buffer.Size() < 4 {
				continue
			}
			bytes := make([]byte, 4)
			err = z.buffer.ReadExactly(bytes)
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			length := int(bytes[3]) | int(bytes[2])<<8 | int(bytes[1])<<16 | int(bytes[0])<<24
			if z.buffer.Size() < length {
				continue
			}
			// in case ddos attack
			if length > z.buffer.Capacity() {
				req.callback(nil, fmt.Errorf("response length %d is too large", length))
				z.closeCh <- struct{}{}
				break
			}
			data := make([]byte, length)
			err = z.buffer.ReadExactly(data)
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			req.callback(data, nil)
		case <-z.closeCh:
			return
		}
	}
}

func (z *ZookeeperClient) write() {
	for {
		select {
		case req := <-z.eventsChan:
			n, err := z.conn.Write(req.bytes)
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			if n != len(req.bytes) {
				req.callback(nil, fmt.Errorf("write %d bytes, but expect %d bytes", n, len(req.bytes)))
				z.closeCh <- struct{}{}
				break
			}
			z.pendingQueue <- req
		case <-z.closeCh:
			return
		}
	}
}

func (z *ZookeeperClient) Close() {
	_ = z.conn.Close()
	z.closeCh <- struct{}{}
}

func NewClient(config ZookeeperClientConfig) (*ZookeeperClient, error) {
	var conn net.Conn
	var err error

	if config.TLSConfig == nil {
		conn, err = net.Dial("tcp", config.addr())
	} else {
		conn, err = tls.Dial("tcp", config.addr(), config.TLSConfig)
	}

	if err != nil {
		return nil, err
	}
	if config.SendQueueSize == 0 {
		config.SendQueueSize = 1000
	}
	if config.PendingQueueSize == 0 {
		config.PendingQueueSize = 1000
	}
	if config.BufferMax == 0 {
		config.BufferMax = 512 * 1024
	}
	z := &ZookeeperClient{}
	z.conn = conn
	z.eventsChan = make(chan *sendRequest, config.SendQueueSize)
	z.pendingQueue = make(chan *sendRequest, config.PendingQueueSize)
	z.buffer = buffer.NewBuffer(config.BufferMax)
	z.closeCh = make(chan struct{})
	go func() {
		z.read()
	}()
	go func() {
		z.write()
	}()
	return z, nil
}
