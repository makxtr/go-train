package main

import "fmt"

func main() {
	// str := "G🧑‍💻o"

	//fmt.Printf("%c ", str[0])

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c ", str[i])
	// }

	// fmt.Println()

	slice := []int{1, 2, 3}

	slice = Append(slice, 4)
	fmt.Println(slice, cap(slice))
}

func Append(slice []int, elem int) []int {
	if len(slice) == cap(slice) {
		newSlice := make([]int, len(slice), cap(slice)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	//for _, e := range elem {
	slice = slice[:len(slice)+1]
	slice[len(slice)-1] = elem
	//}
	return slice
}
