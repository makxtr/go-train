package main

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 10
	fmt.Println("Inside modifyArray", arr)
}

func modifySlice(slice []int) {
	slice[0] = 10
	fmt.Println("Inside modifySlice", slice)
}

func main() {
	array := [3]int{1, 2, 3}
	slice := array[:]

	fmt.Println("Before modifyArray:", array) // 1, 2, 3
	modifyArray(array)                        // 10, 2, 3
	fmt.Println("After modifyArray:", array)  // 1, 2, 3

	fmt.Println("Before modifySlice:", slice) // 1 2 3
	modifySlice(slice)                        // 10 2 3
	fmt.Println("After modifySlice:", slice)  // 10 2 3
	fmt.Println("Final array:", array)        // 10 2 3
}
