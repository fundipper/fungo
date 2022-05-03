package cache

import (
	"fmt"

	set "github.com/deckarep/golang-set/v2"
)

type Set struct{}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Set(key string, value string) bool {
	result, ok := NewCache().Get(s.Key(key))
	if !ok {
		result = set.NewSet[string]()
	}

	ss := result.(set.Set[string])
	if ok := ss.Contains(value); ok {
		return true
	}

	_ = ss.Add(value)
	return NewCache().Set(s.Key(key), result)
}

func (s *Set) Get(key string) ([]string, bool) {
	result, ok := NewCache().Get(s.Key(key))
	if !ok {
		return nil, ok
	}

	ss := result.(set.Set[string])
	return ss.ToSlice(), ok
}

func (s *Set) Key(key string) string {
	return fmt.Sprintf("%s:%s", _SET, key)
}
