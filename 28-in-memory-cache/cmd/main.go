package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type Cache interface {
	Set(k, v string) error
	Get(k string) (string, bool)
}

type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

type MyCache struct {
	shards []*Shard
}

func NewMyCache(shardCount int64) *MyCache {
	shards := make([]*Shard, shardCount)

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}

	return &MyCache{
		shards: shards,
	}
}

func (c *MyCache) Set(k, v string) error {
	shard := c.getShard(k)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.data[k] = v
	return nil
}

func (c *MyCache) Get(k string) (string, bool) {
	shard := c.getShard(k)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	v, ok := shard.data[k]
	return v, ok
}

func (c *MyCache) getShard(k string) *Shard {
	hasher := fnv.New32a()
	hasher.Write([]byte(k))
	hash := hasher.Sum32()
	return c.shards[hash%uint32(len(c.shards))]
}

func main() {
	cache := NewMyCache(100)
	cache.Set("salary", "8000")
	value, ok := cache.Get("salary")
	if !ok {
		fmt.Println("Salary not found")
		return
	}

	fmt.Println("Salary found", value)
}
