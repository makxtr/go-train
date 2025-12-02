package main

import (
	"context"

	"fmt"

	"os"

	"os/signal"

	"sync"
)

type Counter chan int

func NewCounter(ctx context.Context) Counter {
	counter := make(Counter)
	go func() {
		var count int
		for {
			select {
			case v := <-counter:
				count += v
			case counter <- count:
			case <-ctx.Done():
				return
			}
		}
	}()
	return counter
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	counter := NewCounter(ctx)
	var wg sync.WaitGroup

	const total = 1_000_000

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			counter <- 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(<-counter)
}
