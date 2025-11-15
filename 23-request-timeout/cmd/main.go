package main

import (
	"context"
	"fmt"
	"time"
)

func getDiscount() float64 {
	time.Sleep(2 * time.Second)
	return 12.0
}

func getDiscountWrap(ctx context.Context) (float64, error) {
	ch := make(chan float64, 1)

	go func() {
		ch <- getDiscount()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case v := <-ch:
		return v, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := getDiscountWrap(ctx)

	if err != nil {
		fmt.Printf("Error getting discount: %v\n", err)
	} else {
		fmt.Printf("Discount is %f\n", res)
	}
}
