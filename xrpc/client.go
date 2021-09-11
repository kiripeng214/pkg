package xrpc

import (
	"io"
	"kiripeng214/pkg/xrpc/codec"
	"kiripeng214/pkg/xrpc/err"
	"sync"
)

var _ io.Closer = (*Client)(nil)

// Call represents an active RPC.
type Call struct {
	Seq           uint64
	ServiceMethod string // format <service>,<method>
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call
}

func (c *Call) done() {
	c.Done <- c
}

type Client struct {
	cc       codec.Codec
	opt      *Option
	sending  sync.Mutex
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool // user has called Close
	shutdown bool // server has told us to stop
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closing {
		return err.ErrShutdown
	}
	c.closing = true
	return c.cc.Close()
}

func (c *Client) IsAvailable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return !c.shutdown && !c.closing
}
