package cache

import (
	"fmt"
)

type List struct{}

func NewList() *List {
	return &List{}
}

func (l *List) Set(key, value string) bool {
	result, ok := l.Get(key)
	if !ok {
		result = []string{}
	}

	result = append(result, value)
	return NewCache().Set(l.Key(key), result)
}

func (l *List) Get(key string) ([]string, bool) {
	result, ok := NewCache().Get(l.Key(key))
	if !ok {
		return nil, ok
	}
	return result.([]string), ok
}

func (l *List) Key(key string) string {
	return fmt.Sprintf("%s:%s", _LIST, key)
}
