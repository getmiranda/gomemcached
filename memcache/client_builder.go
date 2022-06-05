package memcache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/getmiranda/gomemcached/memcachemock"
)

var (
	DefaultTimeout      = time.Duration(time.Second * 3)
	DefaultMaxIdleConns = 100
)

// ClientBuilder is the interface for building a client.
type ClientBuilder interface {
	// SetTimeout specifies the socket read/write timeout.
	// If zero, DefaultTimeout is used.
	SetTimeout(timeout time.Duration) ClientBuilder
	// MaxIdleConns specifies the maximum number of idle connections that will
	// be maintained per address. If less than one, DefaultMaxIdleConns will be
	// used.
	//
	// Consider your expected traffic rates and latency carefully. This should
	// be set to a number higher than your peak parallel requests.
	SetMaxIdleConns(i int) ClientBuilder
	// WithServers configures the client to use the provided server(s)
	// with equal weight. If a server is listed multiple times,
	// it gets a proportional amount of weight.
	WithServers(servers ...string) ClientBuilder
	// Build builds the memcache client.
	Build() Client
}

type clientBuilder struct {
	timeout      time.Duration
	maxIdleConns int
	servers      []string
}

// SetTimeout specifies the socket read/write timeout.
// If zero, DefaultTimeout is used.
func (c *clientBuilder) SetTimeout(timeout time.Duration) ClientBuilder {
	c.timeout = timeout
	return c
}

// SetMaxIdleConns specifies the maximum number of idle connections that will
// be maintained per address. If less than one, DefaultMaxIdleConns will be
// used.
//
// Consider your expected traffic rates and latency carefully. This should
// be set to a number higher than your peak parallel requests.
func (c *clientBuilder) SetMaxIdleConns(i int) ClientBuilder {
	c.maxIdleConns = i
	return c
}

// WithServers configures the client to use the provided server(s)
// with equal weight. If a server is listed multiple times,
// it gets a proportional amount of weight.
func (c *clientBuilder) WithServers(servers ...string) ClientBuilder {
	c.servers = servers
	return c
}

// Build builds the memcache client.
func (c *clientBuilder) Build() Client {
	if memcachemock.MockupServer.IsEnabled() {
		return memcachemock.MockupServer.GetMockedClient()
	}

	cli := memcache.New(c.servers...)
	cli.Timeout = c.getTimeout()
	cli.MaxIdleConns = c.getMaxIdleConns()

	return &client{cli}
}

func (c *clientBuilder) getTimeout() time.Duration {
	if c.timeout == 0 {
		return DefaultTimeout
	}
	return c.timeout
}

func (c *clientBuilder) getMaxIdleConns() int {
	if c.maxIdleConns == 0 {
		return DefaultMaxIdleConns
	}
	return c.maxIdleConns
}

// NewBuilder creates a new client builder.
func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}
