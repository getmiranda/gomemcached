# Golang Memcached Client

Provides a client for the memcached cache server. It is based on the [memcached](https://pkg.go.dev/github.com/bradfitz/gomemcache) package.

## Prerequisites

In order to use this client you need to have a memcached server running. You can run the following command to start a memcached server with docker:

```bash
docker run -d -p 11211:11211 --name memcached memcached
```

## Instalation

```bash
go get github.com/getmiranda/gomemcached
```

## Usage

In order to use the library you need to import the corresponding package:

```go
package main

import (
    "github.com/getmiranda/gomemcached/memcache"
)
```

### Configuring the client

Once you have imported the package, you can now start using the client. First you need to configure and build the client as you need:

```go
// Create a new builder:
memcacheClient := memcache.NewBuilder().
    // Set the host(s):
    WithServers("localhost:11211").
    // Set the timeout:
    SetTimeout(time.Second * 5).
    // Set the max number of connections per server:
    SetMaxIdleConns(10).
    // Finally, build the client and start using it!
    Build()
```

### Using the client

The `Client` interface provides convenient methods that you can use to perform different operations. For example, you can get a value from the cache:

```go
// Get a value from the cache:
value, err := memcacheClient.Get("key")
```

## Testing

The library provides a convenient package for mocking items and getting a particular values. The mock key is the item key. Every item with the same kay will return the same item mock.

In order to use the mocking features you need to import the corresponding package:

```go
package main

import (
    "github.com/getmiranda/gomemcached/memcachemock"
)
```

### Starting the mock server

```go
func TestMain(m *testing.M) {
    // Tell the library to mock any further operations from here.
    memcachemock.MockupServer.Start()

    // Start the test cases for this pacakge:
    os.Exit(m.Run())
}
```

Once you start the mock server, every operation will be handled by this server and will not be sent against the real memcached server. If there is no mock matching the current operation you'll get an error saying `mock not found`.

### Configuring a given mock

```go
func TestMyTest(t *testing.T) {
    // Delete all mocks in every new test case to ensure a clean environment:
    memcachemock.MockupServer.DeleteMocks()

    // Configure a new mock:
    memcachemock.MockupServer.AddMock(&memcachemock.Mock{
        Operation: memcachemock.OperationGet,
        Args:      memcachemock.Args{"mykey"},

        Error: errors.New("mirandas"),
    })

    it, err := memcacheClient.Get("mykey")

    ...
}
```

In this case, we're telling the memcache client that when we does a get operation with `mykey` argument, we want that particular error. In this case, no result was returned.

Let's see how you can configure a particular value:

```go
func TestMyTest(t *testing.T) {
    // Delete all mocks in every new test case to ensure a clean environment:
    memcachemock.MockupServer.DeleteMocks()

    // Configure a new mock:
    memcachemock.MockupServer.AddMock(&memcachemock.Mock{
        Operation: memcachemock.OperationGet,
        Args:      memcachemock.Args{"mykey"},

        Return: &item.Item{
            Key:   "mykey",
            Value: []byte("myvalue"),
        },
    })

    it, err := memcacheClient.Get("mykey")

    ...
}
```

In this case, we get an item from the cache.
