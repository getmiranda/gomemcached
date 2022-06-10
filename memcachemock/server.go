package memcachemock

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"strings"
	"sync"

	"github.com/getmiranda/gomemcached/item"
)

var (
	MockupServer = mockServer{
		mocks:          make(map[string]*Mock),
		memcacheClient: &clientMock{},
	}

	ErrMockNotFound        = errors.New("mock not found")
	ErrInterfaceConvertion = errors.New("interface convertion error")
)

type mockServer struct {
	enabled        bool
	serverMutex    sync.Mutex
	mocks          map[string]*Mock
	memcacheClient *clientMock
}

// Start sets the enviroment to send all client requests
// to the mockup server.
func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

// Stop stop sending requests to the mockup server
func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

// AddMockups add a mock to the mockup server.
func (m *mockServer) AddMock(mock *Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	if mock == nil {
		return
	}

	m.mocks[m.getMockKey(mock.Operation, mock.Args)] = mock
}

// IsEnabled check whether the mock environment is enabled or not.
func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

// GetMockedClient gets the memcached mock client.
func (m *mockServer) GetMockedClient() *clientMock {
	return m.memcacheClient
}

// DeleteMocks delete all mocks in every new test case to ensure a clean environment.
func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

func (m *mockServer) cleanValue(value []byte) string {
	valueString := string(value)
	valueString = strings.TrimSpace(valueString)
	if valueString == "" {
		return ""
	}
	valueString = strings.ReplaceAll(valueString, "\n", "")
	valueString = strings.ReplaceAll(valueString, "\t", "")
	return valueString
}

func (m *mockServer) getMockKey(op Operation, args ...Args) string {
	key := string(op)
	if len(args) > 0 {
		for _, v := range args[0] {
			var buffer bytes.Buffer
			enc := gob.NewEncoder(&buffer)
			it, ok := v.(*item.Item)
			if ok {
				it.Expiration = 0
				it.Flags = 0
			}
			enc.Encode(v)
			key += m.cleanValue(buffer.Bytes())
		}
	}
	hasher := md5.New()
	hasher.Write([]byte(key))
	newKey := hex.EncodeToString(hasher.Sum(nil))
	return newKey
}
