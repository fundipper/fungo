package cache

import (
	"fmt"
)

type Hash struct{}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) Set(key string, value interface{}) bool {
	return NewCache().Set(h.Key(key), value)
}

func (h *Hash) Get(key string) (interface{}, bool) {
	result, ok := NewCache().Get(h.Key(key))
	if !ok {
		return nil, ok
	}
	return result, ok
}

func (h *Hash) Key(key string) string {
	return fmt.Sprintf("%s:%s", _HASH, key)
}
