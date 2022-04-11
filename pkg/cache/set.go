package cache

import (
	"fmt"
	"sync"
)

type Set struct {
	sync.RWMutex
	store map[string]struct{}
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Push(key string, value string) bool {
	result, ok := s.Get(key)
	if !ok {
		result = map[string]struct{}{}
	}

	set.Lock()
	set.store = result
	set.store[value] = struct{}{}
	set.Unlock()

	return s.Set(key, set.store)
}

func (s *Set) Set(key string, value map[string]struct{}) bool {
	return NewCache().Set(s.Key(key), value)
}

func (s *Set) Get(key string) (map[string]struct{}, bool) {
	data, ok := NewCache().Get(s.Key(key))
	if !ok {
		return nil, ok
	}
	return data.(map[string]struct{}), ok
}

func (s *Set) Key(key string) string {
	return fmt.Sprintf("%s:%s", _SET, key)
}
