package memcache

import (
	"testing"
	"time"

	"github.com/getmiranda/gomemcached/memcachemock"
)

func TestBuilder(t *testing.T) {

	t.Run("NewBuilder", func(t *testing.T) {
		builder := NewBuilder()
		if builder == nil {
			t.Error("NewBuilder returned nil")
		}
	})

	t.Run("SetTimeout", func(t *testing.T) {
		builder := clientBuilder{}
		builder.SetTimeout(time.Second * 5)
		if builder.timeout != time.Second*5 {
			t.Errorf("Expected timeout to be %v, got %v", time.Second*5, builder.timeout)
		}
	})

	t.Run("SetMaxIdleConns", func(t *testing.T) {
		builder := clientBuilder{}
		builder.SetMaxIdleConns(10)
		if builder.maxIdleConns != 10 {
			t.Errorf("Expected maxIdleConns to be %v, got %v", 10, builder.maxIdleConns)
		}
	})

	t.Run("WithServers", func(t *testing.T) {
		builder := clientBuilder{}
		builder.WithServers("localhost:11211")
		if len(builder.servers) != 1 {
			t.Errorf("Expected servers to be %v, got %v", 1, len(builder.servers))
		}
	})

	t.Run("Build", func(t *testing.T) {
		builder := clientBuilder{}
		client := builder.Build()
		if client == nil {
			t.Errorf("Expected client to be non-nil")
		}
	})

	t.Run("BuildGetMockClient", func(t *testing.T) {
		builder := clientBuilder{}
		memcachemock.MockupServer.Start()

		client := builder.Build()
		if client == nil {
			t.Errorf("Expected client to be non-nil")
		}
	})

	t.Run("getTimeoutDefault", func(t *testing.T) {
		builder := clientBuilder{}
		timeout := builder.getTimeout()
		if timeout != time.Second*3 {
			t.Errorf("Expected timeout to be %v, got %v", time.Second*3, timeout)
		}
	})

	t.Run("getTimeoutSet", func(t *testing.T) {
		builder := clientBuilder{}
		builder.SetTimeout(time.Second * 5)
		timeout := builder.getTimeout()
		if timeout != time.Second*5 {
			t.Errorf("Expected timeout to be %v, got %v", time.Second*5, timeout)
		}
	})

	t.Run("getMaxIdleConnsDefault", func(t *testing.T) {
		builder := clientBuilder{}
		maxIdleConns := builder.getMaxIdleConns()
		if maxIdleConns != 100 {
			t.Errorf("Expected maxIdleConns to be %v, got %v", 100, maxIdleConns)
		}
	})

	t.Run("getMaxIdleConnsSet", func(t *testing.T) {
		builder := clientBuilder{}
		builder.SetMaxIdleConns(10)
		maxIdleConns := builder.getMaxIdleConns()
		if maxIdleConns != 10 {
			t.Errorf("Expected maxIdleConns to be %v, got %v", 10, maxIdleConns)
		}
	})
}
