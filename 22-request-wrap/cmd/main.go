package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func SimulateRequest(ctx context.Context) (int64, error) {
	select {
	case <-time.After(time.Duration(rand.Int63n(5)) * time.Second):
		return counter.Add(1), nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	var wg sync.WaitGroup

	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			val, err := SimulateRequest(ctx)
			if err != nil {
				log.Println("Request timed out")
			} else {
				log.Printf("Значение счетчика: %d\n", val)
			}
		}()
	}

	wg.Wait()

	log.Printf("Total counter: %d\n", counter.Load())
}
