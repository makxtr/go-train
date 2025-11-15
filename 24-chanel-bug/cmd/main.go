package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Отдельная горутина отвисла")
		ch <- false
		close(ch)
	}()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Произошел тик тикера")

		case value := <-ch:
			fmt.Printf("Получено значение %t\n", value)
			return
		}
	}
}
