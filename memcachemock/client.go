package memcachemock

import (
	"github.com/getmiranda/gomemcached/item"
)

type clientMock struct{}

func (c *clientMock) FlushAll() error {
	key := MockupServer.getMockKey(OperationFlushAll)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Get(key string) (*item.Item, error) {
	args := Args{key}
	mockKey := MockupServer.getMockKey(OperationGet, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return nil, ErrMockNotFound
	}
	if mock.Error != nil {
		return nil, mock.Error
	}
	return mock.Return.(*item.Item), nil
}

func (c *clientMock) Touch(key string, seconds int32) (err error) {
	args := Args{key, seconds}
	mockKey := MockupServer.getMockKey(OperationTouch, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) GetMulti(keys []string) (map[string]*item.Item, error) {
	args := Args{keys}
	mockKey := MockupServer.getMockKey(OperationGetMulti, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return nil, ErrMockNotFound
	}
	if mock.Error != nil {
		return nil, mock.Error
	}
	return mock.Return.(map[string]*item.Item), nil
}

func (c *clientMock) Set(item *item.Item) error {
	args := Args{item}
	key := MockupServer.getMockKey(OperationSet, args)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Add(item *item.Item) error {
	args := Args{item}
	key := MockupServer.getMockKey(OperationAdd, args)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Replace(item *item.Item) error {
	args := Args{item}
	key := MockupServer.getMockKey(OperationReplace, args)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) CompareAndSwap(item *item.Item) error {
	args := Args{item}
	key := MockupServer.getMockKey(OperationCompareAndSwap, args)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Delete(key string) error {
	args := Args{key}
	mockKey := MockupServer.getMockKey(OperationDelete, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) DeleteAll() error {
	key := MockupServer.getMockKey(OperationDeleteAll)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Ping() error {
	key := MockupServer.getMockKey(OperationPing)
	mock := MockupServer.mocks[key]
	if mock == nil {
		return ErrMockNotFound
	}
	if mock.Error != nil {
		return mock.Error
	}
	return nil
}

func (c *clientMock) Increment(key string, delta uint64) (newValue uint64, err error) {
	args := Args{key, delta}
	mockKey := MockupServer.getMockKey(OperationIncrement, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return 0, ErrMockNotFound
	}
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.Return.(uint64), nil
}

func (c *clientMock) Decrement(key string, delta uint64) (newValue uint64, err error) {
	args := Args{key, delta}
	mockKey := MockupServer.getMockKey(OperationDecrement, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return 0, ErrMockNotFound
	}
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.Return.(uint64), nil
}

func (c *clientMock) Exists(key string) (bool, error) {
	args := Args{key}
	mockKey := MockupServer.getMockKey(OperationExists, args)
	mock := MockupServer.mocks[mockKey]
	if mock == nil {
		return false, ErrMockNotFound
	}
	if mock.Error != nil {
		return false, mock.Error
	}
	return mock.Return.(bool), nil
}
