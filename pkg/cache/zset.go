package cache

import (
	"fmt"
	"sync"

	zset "github.com/yddeng/sortedset"
)

type Zset struct {
	set *zset.SortedSet
	sync.RWMutex
}

func NewZset() *Zset {
	return &Zset{
		set: zset.New(),
	}
}

func (z *Zset) Set(key string, value string, score int64) (int, bool) {
	result, ok := NewCache().Get(z.Key(key))
	if !ok {
		result = NewZset()
	}

	zs := result.(*Zset)
	zs.Lock()
	defer zs.Unlock()

	_ = zs.set.Set(zset.Key(value), Score(score))
	return zs.set.Len(), NewCache().Set(z.Key(key), zs)
}

func (z *Zset) Get(key string, start, end int) ([]string, bool) {
	data, ok := NewCache().Get(z.Key(key))
	if !ok {
		return nil, ok
	}

	zs := data.(*Zset)
	zs.RLock()
	defer zs.RUnlock()

	result := []string{}
	zs.set.Range(start, end, func(rank int, key zset.Key, value interface{}) bool {
		result = append(result, string(key))
		return true
	})
	return result, true
}

func (z *Zset) Len(key string) (int, bool) {
	data, ok := NewCache().Get(z.Key(key))
	if !ok {
		return 0, ok
	}

	zs := data.(*Zset)
	zs.RLock()
	defer zs.RUnlock()

	return zs.set.Len(), true
}

func (z *Zset) Rank(key, value string) (int, bool) {
	data, ok := NewCache().Get(z.Key(key))
	if !ok {
		return 0, ok
	}

	zs := data.(*Zset)
	zs.RLock()
	defer zs.RUnlock()

	return zs.set.GetRank(zset.Key(value)), true
}

func (z *Zset) Key(key string) string {
	return fmt.Sprintf("%s:%s", _ZSET, key)
}

type Score int64

func (this Score) Less(other interface{}) bool {
	return this >= other.(Score)
}
