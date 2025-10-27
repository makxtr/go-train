package main

import "fmt"

func modifyElement(slice []int) {
	slice[1] = 999
}

func addElement(slice []int) {
	slice = append(slice, 100)
	slice[0] = 888
	fmt.Println("Внутри addElement", slice) //  888, 999, 30, 100
}

func main() {
	original := []int{10, 20, 30}

	fmt.Println("До modifyElement", original) // 10,20,30
	modifyElement(original)
	fmt.Println("После modifyElement", original) //10,999,30

	fmt.Println("До addElement", original) //10,999,30
	addElement(original)
	fmt.Println("После addElement", original) //10,999,30
}
