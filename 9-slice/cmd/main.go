package main

import "fmt"

func main() {
	data := []int{10, 20, 30, 40}
	fmt.Println("Изначальный слайс:", data) // 10, 20, 30, 40

	modify(data[:2])
	fmt.Println("Слайс после измененний", data) // 10, 20, 50, 60
}

func modify(slice []int) {
	slice = append(slice, 50, 60)
	fmt.Println("Слайс в функции модификации:", slice) // 10, 20, 50, 60
}
