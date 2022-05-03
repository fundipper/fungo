package cache

import (
	"fmt"
)

type String struct{}

func NewString() *String {
	return &String{}
}

func (s *String) Set(key, value string) bool {
	return NewCache().Set(s.Key(key), value)
}

func (s *String) Get(key string) (string, bool) {
	result, ok := NewCache().Get(s.Key(key))
	if !ok {
		return "", ok
	}
	return result.(string), ok
}

func (s *String) Key(key string) string {
	return fmt.Sprintf("%s:%s", _STRING, key)
}
