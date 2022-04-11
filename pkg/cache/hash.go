package cache

import (
	"fmt"
)

type Hash struct{}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) Set(key string, value map[string]interface{}) bool {
	return NewCache().Set(h.Key(key), value)
}

func (h *Hash) Get(key string) (map[string]interface{}, bool) {
	data, ok := NewCache().Get(h.Key(key))
	if !ok {
		return nil, ok
	}
	return data.(map[string]interface{}), ok
}

func (h *Hash) Key(key string) string {
	return fmt.Sprintf("%s:%s", _HASH, key)
}
