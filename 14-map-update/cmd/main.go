package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]string
	lock sync.RWMutex
}

func (cm *ConcurrentMap) GetOrCreate(key, value string) string {
	cm.lock.RLock()
	val, ok := cm.data[key]
	cm.lock.RUnlock()

	if ok {
		return val
	}

	cm.lock.Lock()
	defer cm.lock.Unlock()

	if val, ok := cm.data[key]; ok {
		return val
	}

	cm.data[key] = value
	return value
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string),
	}
}

func main() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value1")
		fmt.Println("Goroutine 1 go:", val)
	}()

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value2")
		fmt.Println("Goroutine 2 go:", val)
	}()

	wg.Wait()
}
