package main

import "fmt"

type SomeStruct struct {
	Value int
}

func CheckForNil(i interface{}) {
	if i == nil {
		fmt.Println("Nil")
		return
	}

	fmt.Printf("Not nil!")
}

func main() {
	var s *SomeStruct
	CheckForNil(s)
}
