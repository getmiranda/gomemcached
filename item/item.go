package item

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Item memcache.Item

func (item *Item) String() string {
	return fmt.Sprintf("Item{Key: %s, Value: %s}", item.Key, string(item.Value))
}
