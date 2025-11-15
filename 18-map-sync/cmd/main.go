package main

import (
	"fmt"
	"sync"
	"time"
)

import "golang.org/x/sync/singleflight"

type Cache struct {
	m sync.Map
	//mu sync.Mutex
	sf singleflight.Group
}

func (c *Cache) Set(key string, val string) {
	c.m.Store(key, val)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.m.Load(key)
}

func (c *Cache) GetOrCompute(key string, compute func() string) interface{} {
	val, _, _ := c.sf.Do(key, func() (interface{}, error) {
		if v, ok := c.m.Load(key); ok {
			return v, nil
		}
		result := compute()
		c.m.Store(key, result)
		return result, nil
	})
	return val
}

func main() {
	cache := new(Cache)
	var wg sync.WaitGroup

	computeCount := 0

	// Много горутин одновременно запрашивают один ключ
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			result := cache.GetOrCompute("shared_key", func() string {
				computeCount++
				time.Sleep(500 * time.Millisecond) // Долгие вычисления
				return fmt.Sprintf("result_from_goroutine_%d", id)
			})

			fmt.Printf("Горутина %d получила: %v\n", id, result)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Функция compute вызвана %d раз(а)\n", computeCount) // Должно быть 1
}
