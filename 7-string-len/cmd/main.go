package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ddЯЙ描"

	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
