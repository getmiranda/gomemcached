// Package memcached provides a client for the memcached cache server.
package memcache

import (
	"errors"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/getmiranda/gomemcached/item"
)

type Client interface {
	// FlushAll deletes all items in the cache.
	FlushAll() error
	// Get gets the item for the given key. ErrCacheMiss is returned for a
	// memcache cache miss. The key must be at most 250 bytes in length.
	Get(key string) (*item.Item, error)
	// Touch updates the expiry for the given key. The seconds parameter is either
	// a Unix timestamp or, if seconds is less than 1 month, the number of seconds
	// into the future at which time the item will expire. Zero means the item has
	// no expiration time. ErrCacheMiss is returned if the key is not in the cache.
	// The key must be at most 250 bytes in length.
	Touch(key string, seconds int32) (err error)
	// GetMulti is a batch version of Get. The returned map from keys to
	// items may have fewer elements than the input slice, due to memcache
	// cache misses. Each key must be at most 250 bytes in length.
	// If no error is returned, the returned map will also be non-nil.
	GetMulti(keys []string) (map[string]*item.Item, error)
	// Set writes the given item, unconditionally.
	Set(item *item.Item) error
	// Add writes the given item, if no value already exists for its
	// key. ErrNotStored is returned if that condition is not met.
	Add(item *item.Item) error
	// Replace writes the given item, but only if the server *does*
	// already hold data for this key.
	Replace(item *item.Item) error
	// CompareAndSwap writes the given item that was previously returned
	// by Get, if the value was neither modified or evicted between the
	// Get and the CompareAndSwap calls. The item's Key should not change
	// between calls but all other item fields may differ. ErrCASConflict
	// is returned if the value was modified in between the
	// calls. ErrNotStored is returned if the value was evicted in between
	// the calls.
	CompareAndSwap(item *item.Item) error
	// Delete deletes the item with the provided key. The error ErrCacheMiss is
	// returned if the item didn't already exist in the cache.
	Delete(key string) error
	// DeleteAll deletes all items in the cache.
	DeleteAll() error
	// Ping checks all instances if they are alive. Returns error if any
	// of them is down.
	Ping() error
	// Increment atomically increments key by delta. The return value is
	// the new value after being incremented or an error. If the value
	// didn't exist in memcached the error is ErrCacheMiss. The value in
	// memcached must be an decimal number, or an error will be returned.
	// On 64-bit overflow, the new value wraps around.
	Increment(key string, delta uint64) (newValue uint64, err error)
	// Decrement atomically decrements key by delta. The return value is
	// the new value after being decremented or an error. If the value
	// didn't exist in memcached the error is ErrCacheMiss. The value in
	// memcached must be an decimal number, or an error will be returned.
	// On underflow, the new value is capped at zero and does not wrap
	// around.
	Decrement(key string, delta uint64) (newValue uint64, err error)
	// Exists returns true if an item with the given key exists.
	Exists(key string) (bool, error)
}

type client struct {
	mcClient *memcache.Client
}

/* Implementations */

// FlushAll deletes all items in the cache.
func (c *client) FlushAll() error {
	return c.mcClient.FlushAll()
}

// Get gets the item for the given key. ErrCacheMiss is returned for a
// memcache cache miss. The key must be at most 250 bytes in length.
func (c *client) Get(key string) (*item.Item, error) {
	it, err := c.mcClient.Get(key)
	if err != nil {
		return nil, err
	}
	alias := (*item.Item)(it)
	return alias, nil

}

// Touch updates the expiry for the given key. The seconds parameter is either
// a Unix timestamp or, if seconds is less than 1 month, the number of seconds
// into the future at which time the item will expire. Zero means the item has
// no expiration time. ErrCacheMiss is returned if the key is not in the cache.
// The key must be at most 250 bytes in length.
func (c *client) Touch(key string, seconds int32) (err error) {
	return c.mcClient.Touch(key, seconds)
}

// GetMulti is a batch version of Get. The returned map from keys to
// items may have fewer elements than the input slice, due to memcache
// cache misses. Each key must be at most 250 bytes in length.
// If no error is returned, the returned map will also be non-nil.
func (c *client) GetMulti(keys []string) (map[string]*item.Item, error) {
	multi, err := c.mcClient.GetMulti(keys)
	if err != nil {
		return nil, err
	}
	items := make(map[string]*item.Item)
	for k, v := range multi {
		items[k] = (*item.Item)(v)
	}
	return items, nil
}

// Set writes the given item, unconditionally.
func (c *client) Set(item *item.Item) error {
	alias := (*memcache.Item)(item)
	return c.mcClient.Set(alias)
}

// Add writes the given item, if no value already exists for its
// key. ErrNotStored is returned if that condition is not met.
func (c *client) Add(item *item.Item) error {
	alias := (*memcache.Item)(item)
	return c.mcClient.Add(alias)
}

// Replace writes the given item, but only if the server *does*
// already hold data for this key.
func (c *client) Replace(item *item.Item) error {
	alias := (*memcache.Item)(item)
	return c.mcClient.Replace(alias)
}

// CompareAndSwap writes the given item that was previously returned
// by Get, if the value was neither modified or evicted between the
// Get and the CompareAndSwap calls. The item's Key should not change
// between calls but all other item fields may differ. ErrCASConflict
// is returned if the value was modified in between the
// calls. ErrNotStored is returned if the value was evicted in between
// the calls.
func (c *client) CompareAndSwap(item *item.Item) error {
	alias := (*memcache.Item)(item)
	return c.mcClient.CompareAndSwap(alias)
}

// Delete deletes the item with the provided key. The error ErrCacheMiss is
// returned if the item didn't already exist in the cache.
func (c *client) Delete(key string) error {
	return c.mcClient.Delete(key)
}

// DeleteAll deletes all items in the cache.
func (c *client) DeleteAll() error {
	return c.mcClient.DeleteAll()
}

// Ping checks all instances if they are alive. Returns error if any
// of them is down.
func (c *client) Ping() error {
	return c.mcClient.Ping()
}

// Increment atomically increments key by delta. The return value is
// the new value after being incremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On 64-bit overflow, the new value wraps around.
func (c *client) Increment(key string, delta uint64) (newValue uint64, err error) {
	return c.mcClient.Increment(key, delta)
}

// Decrement atomically decrements key by delta. The return value is
// the new value after being decremented or an error. If the value
// didn't exist in memcached the error is ErrCacheMiss. The value in
// memcached must be an decimal number, or an error will be returned.
// On underflow, the new value is capped at zero and does not wrap
// around.
func (c *client) Decrement(key string, delta uint64) (newValue uint64, err error) {
	return c.mcClient.Decrement(key, delta)
}

// Exists returns true if an item with the given key exists.
func (c *client) Exists(key string) (bool, error) {
	it, err := c.mcClient.Get(key)
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return false, nil
		}
		return false, err
	}
	return it != nil, nil
}
