package main

import "fmt"

func main() {
	// a1 := []int{1, 2, 3, 4, 5}
	// a2 := append(a1, 6) // 1, 2, 3, 4, 5, 6
	// a3 := append(a1, 7) // 1, 2, 3, 4, 5, 7

	// fmt.Println(a1, a2, a3)

	a1 := make([]int, 0, 10)

	a1 = append(a1, []int{1, 2, 3, 4, 5}...)

	a2 := append(a1, 6) // 1, 2, 3, 4, 5, 6 ptr - 0xb
	a3 := append(a1, 7) // 1, 2, 3, 4, 5, 7 ptr - 0xb

	fmt.Println(a1, a2, a3)
}
