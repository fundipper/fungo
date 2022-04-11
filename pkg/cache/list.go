package cache

import (
	"fmt"
)

type List struct{}

func NewList() *List {
	return &List{}
}
func (l *List) Push(key, value string) bool {
	result, ok := l.Get(key)
	if !ok {
		result = []string{}
	}
	result = append(result, value)
	return l.Set(key, result)
}

func (l *List) Set(key string, value []string) bool {
	return NewCache().Set(l.Key(key), value)
}

func (l *List) Get(key string) ([]string, bool) {
	data, ok := NewCache().Get(l.Key(key))
	if !ok {
		return nil, ok
	}
	return data.([]string), ok
}

func (l *List) Key(key string) string {
	return fmt.Sprintf("%s:%s", _LIST, key)
}
