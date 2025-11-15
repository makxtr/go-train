package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	channels := make([]chan int64, 10000)

	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	for v := range merge(ctx, channels...) {
		println(v)
	}
}

func merge(ctx context.Context, channels ...chan int64) chan int64 {
	res := make(chan int64)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan int64) {
			defer wg.Done()
			// for v := range ch {
			// 	res <- v
			// }

			for {
				select {
				case <-ctx.Done():
					fmt.Println("Context done")
					return
				case val, ok := <-ch:
					if !ok {
						return
					}
					res <- val
				}
			}

		}(ch)
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
