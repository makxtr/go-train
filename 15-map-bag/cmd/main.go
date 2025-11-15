package main

import (
	"fmt"
	"time"
)

func UpdateProductStock() <-chan map[string]int {
	stockUpdates := make(chan map[string]int)

	go func() {
		defer close(stockUpdates)

		currentStock := map[string]int{
			"Apples":  50,
			"Bananas": 30,
			"Oranges": 20,
			"Grapes":  15,
		}

		for i := 0; i < 5; i++ {
			for product, quantity := range currentStock {
				currentStock[product] = int(float64(quantity) * 0.95)
			}

			stock := make(map[string]int, len(currentStock))
			for product, quantity := range currentStock {
				stock[product] = quantity
			}

			stockUpdates <- stock

			time.Sleep(150 * time.Millisecond)
		}
	}()

	return stockUpdates
}

func main() {
	stockStream := UpdateProductStock()

	stockHistory := make([]map[string]int, len(stockStream))

	for stock := range stockStream {
		stockHistory = append(stockHistory, stock)
	}

	for i, stock := range stockHistory {
		fmt.Printf("Iteration %d: %+v\n", i+1, stock)
	}
}
